package fx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/getsentry/raven-go"
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
	opts          *options
}

func NewEthExecutor(keystore *keychain.Store, conf g.BlockChainConfig, db *sql.DB, eos ...ExecuteOption) (*EthExecutor, error) {
	executor := new(EthExecutor)
	executor.keystore = keystore
	executor.ethUrl = conf.RawUrl
	executor.contractAddrs = conf.ContractAddrs
	executor.db = db

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
	c, err1 := clientCache.GetOrCreate(companyID)
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
	p := &CmdProcessor{fxClient: adminClient, cmd: cmd, db: e.db}
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

func (e *EthExecutor) pay(p *CmdProcessor) error {
	cmd := p.cmd
	output := cmd.Tx.Output
	input := cmd.Tx.Input

	var tx *ethTypes.Transaction
	for idx, t := range output {
		boxId := generateBoxId(cmd.Tx.TxId, idx)
		to := t.Owner
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

		r, err := p.WaitMined(context.Background(), tx)
		if err != nil {
			log.Errorf("create payBox(id: %v) failed: %v", boxId, err)
			return err
		}

		n, err := e.boxing(p, t.Amount, input, r.ContractAddress, boxId)
		if err != nil {
			log.Errorf("transfer to box(id: %v) failed: %v", boxId, err)
			return err
		}

		log.Infof("transfer box(%v) to %v", boxId, t.Owner)
		acc := e.keystore.GetAdminAccount()
		box, err := contract_gen.NewFuxPayBox(r.ContractAddress, p.fxClient.EthClient())
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
	var tx *ethTypes.Transaction
	size := len(tokens)
	if size > batchSize {
		size = batchSize
	}
	tokenIds := make([]*big.Int, size)

	packing := func() error {
		jobId := new(big.Int).SetUint64(boxId*10 + uint64((consumedNum-1)/batchSize))
		err := p.CallWithFxBatchTransactor(
			func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.AddJob(jobId, fromAcc.Address, boxAddr, tokenIds)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxBatch.AddJob contract failed: %v", err)
			return err
		}

		err = p.CallWithFxBatchTransactor(
			func(session *contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error) {
				var innerErr error
				tx, innerErr = session.RunJob(jobId)
				return tx, innerErr
			},
		)
		if err != nil {
			log.Errorf("call FuxBatch.RunJob contract failed: %v", err)
			return err
		}

		return nil
	}

	for idx, t := range tokens {
		if actualAmount >= targetAmount {
			break
		}

		actualAmount += t.Amount
		consumedNum += 1
		i := idx % batchSize
		tokenIds[i] = &t.Id
		if actualAmount%batchSize == 0 {
			if err := packing(); err != nil {
				return consumedNum, err
			}

			size := len(tokens) - consumedNum
			if size > batchSize {
				size = batchSize
			}
			tokenIds = make([]*big.Int, size)
		}
	}

	if actualAmount%batchSize > 0 {
		if err := packing(); err != nil {
			return consumedNum, err
		}
	}

	return consumedNum, nil
}

func generateBoxId(txId uint64, index int) uint64 {
	return txId*10 + uint64(index)
}
