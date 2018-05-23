package fx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/getsentry/raven-go"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
)

var (
	ErrTxExecuteFailed = errors.New("transaction execute failed")
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

	switch cmd.Tx.TxType {
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
	var err error
	switch cmd.Tx.TxType {
	case SplitFX:
		err = e.splitFX(p)
	case Discount:
		err = e.pay(p)
	case Payment:
		err = e.pay(p)
	case MintFX:
		err = e.mintFX(p)
	case Confirm:
		err = e.confirm(p)
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

	var tokens [2]Token
	copy(tokens[:], cmd.Tx.Output[:2])

	var newTokenIds [2]*big.Int
	var amounts [2]*big.Int
	var states [2]*big.Int
	for i, t := range tokens {
		newTokenIds[i] = &t.Id
		amounts[i] = new(big.Int).SetUint64(t.Amount)
		states[i] = new(big.Int).SetInt64(int64(t.State))
	}
	log.Infof("--- split fx: tokenId: %v, newTokenIds: %+v, amounts: %+v", tokenId.String(), newTokenIds, amounts)

	var tx *ethTypes.Transaction
	err := p.CallWithFxSplitTransactor(
		func(session *contract_gen.FuxSplitTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.SplitFux(&tokenId, newTokenIds, amounts, states)
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

//This pay method with AddJob and RunJob
func (e *EthExecutor) pay(p *CmdProcessor) error {
	cmd := p.cmd
	input := cmd.Tx.Input
	m := mergeToken(cmd.Tx.Output)
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var tx *ethTypes.Transaction
	for idx, owner := range keys {
		boxId := generateBoxId(cmd.Tx.TxId, idx)
		to := owner
		toAcc, err := e.keystore.GetAccount(to)
		if err != nil {
			log.Errorf("get account of %v failed: %v", to, err)
			return err
		}
		log.Infof("create box: %v for %s", boxId, to)

		err = p.CallWithBoxFactoryTransactor(
			func(session *contract_gen.FuxPayBoxFactoryTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.CreatePayBox(new(big.Int).SetUint64(boxId), toAcc.Address)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxPayBoxFactory.CreatePayBox contract failed: %v", err)
			return err
		}

		var boxAddr common.Address
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		r, err := p.WaitMined(ctx, tx)
		cancel()
		if err != nil {
			// Compatible with Ganache backend
			err1 := p.fxClient.CallWithBoxFactoryCaller(
				func(session *contract_gen.FuxPayBoxFactoryCallerSession) error {
					addr, innerErr := session.GetPayBoxAddress(new(big.Int).SetUint64(boxId))
					boxAddr = addr
					return innerErr
				},
			)
			if err1 != nil {
				log.Errorf("create payBox(id: %v) failed: %v", boxId, err)
				return err
			}
		} else {
			boxAddr = r.ContractAddress
		}

		amount := m[owner]
		n, err := e.boxing(p, amount, input, boxAddr, boxId)
		if err != nil {
			log.Errorf("transfer to box(id: %v) failed: %v", boxId, err)
			return err
		}

		log.Infof("transfer box(%v) to %v", boxId, owner)
		acc := e.keystore.GetAdminAccount()
		box, err := contract_gen.NewFuxPayBox(boxAddr, p.fxClient.EthClient())
		if err != nil {
			log.Errorf("Create new FuxPayBox contract failed: %v", err)
			return err
		}

		err = p.CallWithBoxTransactor(
			box,
			func(session *contract_gen.FuxPayBoxTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.TransferOwnership(acc.Address)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxPayBox.TransferOwnership contract failed: %v", err)
			return err
		}

		input = input[n:]
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

//This pay method with batch transfer
func (e *EthExecutor) payByTransfer(p *CmdProcessor) error {
	toAccount := p.cmd.Tx.Output[0].Owner
	Acc, err := e.keystore.GetAccount(toAccount)
	if err != nil {
		log.Errorf("GET Account failed :%v", err)
		return err
	}
	id := make([]*big.Int, len(p.cmd.Tx.Input))
	for i, input := range p.cmd.Tx.Input {
		id[i] = &input.Id
	}

	log.Infof("Transfer FX from:%v to:%v, FX id is :%v", p.cmd.Tx.Input[0].Owner, toAccount, id)
	var tx *ethTypes.Transaction

	err = p.CallWithFxBatchTransactor(
		func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
			var innerErr error
			tx, innerErr = session.Transfer(Acc.Address, id)
			return tx, innerErr
		},
	)

	receipt, err := p.WaitMined(context.Background(), tx)
	if err != nil {
		log.Errorf("transfer failed : %v", err)
		return err
	}

	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return ErrTxExecuteFailed
	}
	log.Infof("transfer success :%v ,receipt info:%v", tx, receipt)
	return nil
}

// Input tokens should be same with pay input tokens
func (e *EthExecutor) confirm(p *CmdProcessor) error {
	cmd := p.cmd
	txId := cmd.Tx.TxId
	boxId := generateBoxId(txId, 0)

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
	output := cmd.Tx.Output
	var tx *ethTypes.Transaction
	for _, t := range output {
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
					new(big.Int).SetInt64(t.ExpireTime),
					new(big.Int).SetUint64(uint64(t.State)),
					"",
				)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxToken.Mint contract failed: %v", err)
			return err
		}
	}
	return nil

}

func (e *EthExecutor) boxing(p *CmdProcessor, targetAmount uint64, tokens []Token, boxAddr common.Address, boxId uint64) (int, error) {
	companyID := p.cmd.Tx.Sponsor()
	fromAcc, err := e.keystore.GetAccount(companyID)
	if err != nil {
		log.Errorf("get account of %v failed: %v", companyID, err)
		return 0, err
	}

	var actualAmount uint64
	var consumedNum int
	tokenIds := make([]*big.Int, len(tokens))
	for idx, t := range tokens {
		if actualAmount >= targetAmount {
			break
		}

		actualAmount += t.Amount
		consumedNum += 1
		tokenIds[idx] = &t.Id
	}

	tokenIds = tokenIds[:consumedNum]
	jobId := new(big.Int).SetUint64(boxId)
	err = p.CallWithFxBatchTransactor(
		func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
			return session.AddJob(jobId, fromAcc.Address, boxAddr, tokenIds)
		},
	)
	if err != nil {
		log.Errorf("call FuxBatch.AddJob contract failed: %v", err)
		return 0, err
	}

	n := consumedNum / batchSize
	if consumedNum%batchSize > 0 {
		n += 1
	}
	for i := 0; i < n; i++ {
		err = p.CallWithFxBatchTransactor(
			func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
				return session.RunJob(jobId)
			},
		)
		if err != nil {
			log.Errorf("call FuxBatch.RunJob contract failed: %v", err)
			return 0, err
		}
	}

	return consumedNum, nil
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

func HandleTransaction(supplierChain chan Transaction, resultChan chan ProcessResult) {
	for transaction := range supplierChain {

		if len(supplierChain) == 0 {
			break
		}
		var ethTransaction *ethTypes.Transaction
		result := &ProcessResult{Id: transaction.Id, Supplier: transaction.Input[0].Owner}
		executor, err := NewEthExecutor(keystore, bConf, db, clientCache, adminClient)
		hash := make(map[string]string)
		cmd := Command{Tx: transaction, txHashes: hash}
		err = executor.Execute(cmd)
		if err != nil {
			log.Errorf("Execute Transaction error:%v", err)
			result.err = err
			resultChan <- *result
		}
		result.Tx = ethTransaction
		resultChan <- *result
	}
}
