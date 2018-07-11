package fx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"math/big"
	"sync"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/getsentry/raven-go"
)

var (
	ErrTxExecuteFailed = errors.New("transaction execute failed")
)

var (
	defaultExecutor *EthExecutor
	executorOnce    sync.Once
)

const (
	batchSize = 30
)

type Executor interface {
	Execute(cmd Command) error
}

type EthExecutor struct {
	ethUrl        string
	contractAddrs g.ContractAddrs
	keystore      *keychain.Store
	db            *sql.DB
	clientCache   *ClientCache
	adminClient   *blockchain.FxClient
	opts          *options
}

func DefaultExecutor() *EthExecutor {
	executorOnce.Do(func() {
		conf := g.GetConfig()
		bConf := conf.BlockchainConfig
		db, err := g.OpenDB(conf.DbConfig)
		if err != nil {
			panic(err)
		}
		store := keychain.DefaultStore()
		cache := DefaultCache()
		adminClient, err := blockchain.NewFxClient(store.GetAdminClient(), store.GetAdminAccount(), bConf.ContractAddrs)
		if err != nil {
			panic(err)
		}
		defaultExecutor, err = NewEthExecutor(store, bConf, db, cache, adminClient)
		if err != nil {
			panic(err)
		}
	})
	return defaultExecutor
}

func NewEthExecutor(
	keystore *keychain.Store,
	conf g.BlockChainConfig,
	db *sql.DB,
	clientCache *ClientCache,
	adminClient *blockchain.FxClient,
	eos ...ExecuteOption,
) (*EthExecutor, error) {
	executor := new(EthExecutor)
	executor.keystore = keystore
	executor.ethUrl = conf.RawUrl
	executor.contractAddrs = conf.ContractAddrs
	executor.db = db
	executor.clientCache = clientCache
	executor.adminClient = adminClient

	executor.opts = defaultOptions
	for _, o := range eos {
		o(executor.opts)
	}

	return executor, nil
}

func (e *EthExecutor) handleError(err error, cmd Command, processor *CmdProcessor) {
	processor.fxClient.RefreshNonce()

	switch err {
	case core.ErrInsufficientFunds:
		companyId := cmd.Tx.Sponsor()
		if acc, err1 := e.keystore.GetAccount(companyId); err1 == nil {
			amount := new(big.Int)
			amount = amount.Mul(keychain.Eth1(), big.NewInt(10))
			e.keystore.TransferEther(acc.Address, amount)
		}
	default:
		raven.CaptureError(err, nil)

	}
}

func (e *EthExecutor) Execute(cmd Command) error {
	log.Infof("execute command: %+v", cmd)
	txType, err := ParseType(cmd.Tx.TxType)
	if err != nil {
		return err
	}
	switch txType {
	case SplitFX, Discount, Payment:
		return e.executeBySupplier(cmd)
	case MintFX, Confirm:
		return e.executeByPlatform(cmd)
	default:
		return fmt.Errorf("unknow command: %+v", cmd)
	}
}

func (e *EthExecutor) executeBySupplier(cmd Command) error {
	companyID := cmd.Tx.Sponsor()
	c, err1 := e.clientCache.GetOrCreate(companyID)
	if err1 != nil {
		log.Errorf("create personal client failed: %v", err1)
		return err1
	}
	p := &CmdProcessor{fxClient: c, cmd: cmd, db: e.db}

	return e.process(cmd, p)
}

func (e *EthExecutor) process(cmd Command, p *CmdProcessor) error {
	if e.opts.max == 0 {
		err := e.run(cmd, p)
		if err != nil {
			e.handleError(err, cmd, p)
		}
		return err
	}

	var lastErr error
	for attempt := uint(0); attempt < e.opts.max; attempt++ {
		waitRetryBackoff(attempt, e.opts)

		lastErr = e.run(cmd, p)
		if lastErr == nil {
			return nil
		}

		e.handleError(lastErr, cmd, p)
		if !e.opts.retryIf(lastErr) {
			return lastErr
		}

		log.Warnf("executor retry attempt: %d, err: %v", attempt+1, lastErr)
	}

	return lastErr
}

func (e *EthExecutor) run(cmd Command, p *CmdProcessor) error {
	txType, err := ParseType(cmd.Tx.TxType)
	if err != nil {
		return err
	}
	switch txType {
	case SplitFX:
		err = e.splitFX(p)
	case Discount:
		err = e.payByTransfer(p)
	case Payment:
		err = e.payTransaction(p)
	case MintFX:
		err = e.mintTransaction(p)
	case Confirm:
		err = e.payByTransfer(p)
	default:
		return fmt.Errorf("err command: %v", cmd)
	}

	if err == nil {
		p.finishProcedure()
	}

	return err
}

func (e *EthExecutor) executeByPlatform(cmd Command) error {
	p := &CmdProcessor{fxClient: e.adminClient, cmd: cmd, db: e.db}
	return e.process(cmd, p)
}

func (e *EthExecutor) splitFX(p *CmdProcessor) error {
	cmd := p.cmd
	token := cmd.Tx.Input[0]
	tokenId := token.Id

	var newTokenIds, amounts, states []*big.Int
	for _, t := range p.cmd.Tx.Output {
		id := t.Id
		newTokenIds = append(newTokenIds, &id)
		amounts = append(amounts, new(big.Int).SetUint64(t.Amount))
		state, err := ParseState(t.State)
		if err != nil {
			return err
		}
		states = append(states, new(big.Int).SetInt64(int64(state)))
	}
	log.Infof("--- split fx: tokenId: %v, newTokenIds: %+v, amounts: %+v", tokenId.String(), newTokenIds, amounts)

	var tx *ethTypes.Transaction
	err := p.CallWithFxSpliterTransactor(
		func(session *contract_gen.FuxSpliterTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.Split(&tokenId, newTokenIds, amounts, states)
			return tx, innerErr
		},
	)
	if err != nil {
		log.Errorf("call FuxSplit.SplitFux contract failed: %v", err)
		return err
	}

	receipt, err := p.WaitMined(context.Background(), tx)
	if err != nil {
		return err
	}
	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return ErrTxExecuteFailed
	}

	return nil
}

func (e *EthExecutor) payTransaction(p *CmdProcessor) error {
	cmd := p.cmd
	input := cmd.Tx.Input
	output := cmd.Tx.Output
	owner := input[0].Owner
	var transferToken []Token

	for n, inputToken := range input {
		var splitToken []Token
		for _, outputToken := range output {
			if inputToken.Id.Uint64() == outputToken.ParentId.Uint64() {
				splitToken = append(splitToken, outputToken)
			}
			if outputToken.Owner != owner && n == 0 {
				transferToken = append(transferToken, outputToken)
			}
		}
		if len(splitToken) == 1 {
			splitToken = append(splitToken, Token{Amount: 0, Id: *big.NewInt(0), State: "Normal"})
		}
		if len(splitToken) == 0 {
			log.Errorf("Invalid transation:%v ,can't find parentId for output", p.cmd.Tx.TxId)
			return fmt.Errorf("Invalid transation:%v ,can't find parentId for output", p.cmd.Tx.TxId)
		}
		err := e.split(inputToken, splitToken, p)
		if err != nil {
			log.Errorf("split fx failed: %v,txid:%v", p.cmd.Tx.TxId, err)
			return err
		}
	}
	err := e.batchTransfer(input[0], transferToken, p)
	if err != nil {
		log.Errorf("transfer fx failed: %v,txid:%v", p.cmd.Tx.TxId, err)
		return err
	}
	return nil
}

//This pay method with batch transfer
func (e *EthExecutor) payByTransfer(p *CmdProcessor) error {
	fromAccount := p.cmd.Tx.Input[0].Owner
	toAccount := p.cmd.Tx.Output[0].Owner
	toAcc, err := e.keystore.GetAccount(toAccount)
	fromAcc, err := e.keystore.GetAccount(fromAccount)
	if err != nil {
		log.Errorf("GET Account failed :%v", err)
		return err
	}

	ids := make([]*big.Int, len(p.cmd.Tx.Input))
	for i, t := range p.cmd.Tx.Input {
		id := t.Id
		ids[i] = &id
		var tx *ethTypes.Transaction
		err = p.CallWithFxTokenTransactor(
			func(session *contract_gen.FuxTokenTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.SafeTransferFrom(fromAcc.Address, toAcc.Address, &id)
				return tx, innerErr
			},
		)

		if err != nil {
			log.Errorf("call FuxBatch.Transfer contract failed: %v", err)
			return err
		}

		receipt, err := p.WaitMined(context.Background(), tx)
		if err != nil {
			log.Errorf("transfer failed : %v", err)
			return err
		}
		if receipt.Status == ethTypes.ReceiptStatusFailed {
			return ErrTxExecuteFailed
		}
		log.Infof("transfer success :%v ,receipt info:%v", tx, receipt.TxHash)
	}
	log.Infof("Transfer FX from:%v to:%v, FX id is :%v", fromAccount, toAccount, ids)

	return nil
}

func (e *EthExecutor) batchTransfer(input Token, output []Token, p *CmdProcessor) error {
	fromAccount := input.Owner
	toAccount := output[0].Owner
	toAcc, err := e.keystore.GetAccount(toAccount)
	fromAcc, err := e.keystore.GetAccount(fromAccount)
	if err != nil {
		log.Errorf("GET Account failed :%v", err)
		return err
	}
	ids := make([]*big.Int, len(output))
	for i, t := range output {
		id := t.Id
		ids[i] = &id
	}
	log.Infof("Transfer FX from:%v to:%v, FX id is :%v", fromAccount, toAccount, ids)
	var tx *ethTypes.Transaction
	err = p.CallWithFxBatchTransactor(
		func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.SafeTransferFrom(fromAcc.Address, toAcc.Address, ids)
			return tx, innerErr
		},
	)

	if err != nil {
		log.Errorf("call FuxBatch.Transfer contract failed: %v", err)
		return err
	}

	receipt, err := p.WaitMined(context.Background(), tx)
	if err != nil {
		log.Errorf("transfer failed : %v", err)
		return err
	}
	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return ErrTxExecuteFailed
	}
	log.Infof("transfer success,input: %v,output:%v", input, output)

	return nil
}

func (e *EthExecutor) mintTransaction(p *CmdProcessor) error {
	err := e.mintFX(p)
	if err != nil {
		log.Errorf("mint fx failed: %v, txid:%v", err, p.cmd.Tx.TxId)
		return err
	}
	cmd := p.cmd
	input := cmd.Tx.Input
	output := cmd.Tx.Output
	for _, t := range input {
		var splitToken []Token
		for _, st := range output {
			if t.Id.Uint64() == st.ParentId.Uint64() {
				splitToken = append(splitToken, st)
			}
		}
		err = e.split(t, splitToken, p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EthExecutor) split(input Token, output []Token, p *CmdProcessor) error {
	tokenId := input.Id
	var newTokenIds, amounts, states []*big.Int
	for _, t := range output {
		id := t.Id
		newTokenIds = append(newTokenIds, &id)
		amounts = append(amounts, new(big.Int).SetUint64(t.Amount))
		state, err := ParseState(t.State)
		if err != nil {
			return err
		}
		states = append(states, new(big.Int).SetInt64(int64(state)))
	}
	log.Infof("Split fx: tokenId: %v, newTokenIds: %+v, amounts: %+v", tokenId.String(), newTokenIds, amounts)
	var tx *ethTypes.Transaction
	err := p.CallWithFxSpliterTransactor(
		func(session *contract_gen.FuxSpliterTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.Split(&tokenId, newTokenIds, amounts, states)
			return tx, innerErr
		},
	)
	if err != nil {
		log.Errorf("call FuxSplit.SplitFux contract failed: %v", err)
		return err
	}

	receipt, err := p.WaitMined(context.Background(), tx)
	if err != nil {
		return err
	}
	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return ErrTxExecuteFailed
	}

	return nil
}

// Input tokens should be same with pay input tokens
func (e *EthExecutor) confirm(p *CmdProcessor) error {
	cmd := p.cmd
	//txId := cmd.Tx.TxId
	txId := 0
	boxId := generateBoxId(uint64(txId), 0)

	var boxAddr common.Address
	err := p.fxClient.CallWithBoxFactoryCaller(
		func(session *contract_gen.FuxPayBoxFactoryCallerSession) error {
			var innerErr error
			boxAddr, innerErr = session.GetPayBoxAddress(new(big.Int).SetUint64(boxId))
			return innerErr
		},
	)
	if err != nil {
		return err
	}
	box, err := contract_gen.NewFuxPayBox(boxAddr, p.fxClient.EthClient())
	if err != nil {
		return err
	}

	input := cmd.Tx.Input
	to := cmd.Tx.Output[0].Owner
	toAcc, err := e.keystore.GetAccount(to)
	if err != nil {
		log.Errorf("get account of %v failed: %v", to, err)
		return err
	}

	var tx *ethTypes.Transaction
	for _, t := range input {
		err := p.CallWithBoxTransactor(
			box,
			func(session *contract_gen.FuxPayBoxTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.Transfer(toAcc.Address, &t.Id)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxPayBox.Transfer contract failed: %v", err)
			return err
		}
	}

	log.Infof("destroy box: %v", boxId)
	err = p.CallWithBoxFactoryTransactor(
		func(session *contract_gen.FuxPayBoxFactoryTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.CloseBox(boxAddr)
			return tx, innerErr
		},
	)
	if err != nil {
		log.Errorf("call FuxPayBoxFactory.CloseBox contract failed: %v", err)
		return err
	}

	receipt, err := p.WaitMined(context.Background(), tx)
	if err != nil {
		return err
	}
	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return ErrTxExecuteFailed
	}

	return nil
}

func (e *EthExecutor) mintFX(p *CmdProcessor) error {
	cmd := p.cmd
	input := cmd.Tx.Input
	var tx *ethTypes.Transaction
	for _, t := range input {
		log.Infof("mint fx: id: %v, owner: %v, amount: %v, state: %v", t.Id.String(), t.Owner, t.Amount, t.State)
		acc, err := e.keystore.GetAccount(t.Owner)
		if err != nil {
			log.Errorf("get account of %v failed: %v", t.Owner, err)
			return err
		}

		err = p.CallWithFxTokenTransactor(
			func(session *contract_gen.FuxTokenTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.Mint(
					acc.Address,
					&t.Id,
					new(big.Int).SetUint64(t.Amount),
					new(big.Int).SetUint64(uint64(Normal)),
					new(big.Int).SetInt64(t.ExpireTime),
				)
				return tx, innerErr
			},
		)

		if err != nil {
			log.Errorf("call FuxToken.Mint contract failed: %v", err)
			return err
		}

		receipt, err := p.WaitMined(context.Background(), tx)
		if err != nil {
			return err
		}
		if receipt.Status == ethTypes.ReceiptStatusFailed {
			return ErrTxExecuteFailed
		}
	}
	return nil
}

func generateBoxId(txId uint64, index int) uint64 {
	return txId*10 + uint64(index)
}

func mergeToken(input []Token) map[string]uint64 {
	m := make(map[string]uint64)
	for _, t := range input {
		owner := t.Owner
		v, ok := m[owner]
		if !ok {
			v = 0
		}
		m[owner] = v + t.Amount
	}
	return m
}

func HandleTransaction(supplierChain chan Transaction) {
	if transaction, ok := <-supplierChain; ok {
		executor := DefaultExecutor()
		hash := make(map[string]string)
		cmd := Command{Tx: transaction, txHashes: hash}
		err := executor.Execute(cmd)
		log.Infof("Complete Transaction，transaction id:%v .", transaction.TxId)
		if err != nil {
			log.Errorf("Execute Transaction error，transaction id:%v , error message:%v", transaction.TxId, err)
		}
	}
	if len(supplierChain) > 0 {
		HandleTransaction(supplierChain)
	}
}
