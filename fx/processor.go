package fx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"encoding/json"
	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrTxExecuteFailed = errors.New("transaction execute failed")
)

const (
	stateProcessing = "processing"
	stateProcessed  = "processed"

	batchSize = 30
)

type CmdProcessor struct {
	fxClient *blockchain.FxClient
	cmd      Command
	db       *sql.DB
}

func (p *CmdProcessor) CallWithFxTokenTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxTokenTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithFxTokenTransactor(ctx, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxSplitTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxSplitTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithFxSplitTransactor(ctx, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxBatchTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxBatchTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithFxBatchTransactor(ctx, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) initCmdNonce() {
	if p.cmd.startNonce == 0 {
		currNonce := p.fxClient.GetNonce()
		p.cmd.startNonce = currNonce
		p.cmd.currNonce = currNonce
		log.Infof(
			"init command procedure and write to db: command_id: %v, start_nonce: %v",
			p.cmd.Tx.Id,
			p.cmd.startNonce,
		)
		p.createProcedure()
	}
}

func (p *CmdProcessor) createProcedure() error {
	stmt, err := p.db.Prepare("INSERT INTO cmd_procedure(command_id, start_nonce, state) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, p.cmd.Tx.Id, p.cmd.startNonce, stateProcessing)
	return err
}

func (p *CmdProcessor) finishProcedure() error {
	stmt, err := p.db.Prepare("UPDATE cmd_procedure SET state = ? WHERE command_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, stateProcessed, p.cmd.Tx.Id)
	return err
}

func (p *CmdProcessor) updateReceipt(receipt *ethTypes.Receipt) error {
	b, err := json.Marshal(receipt)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		"UPDATE cmd_procedure set receipts = JSON_SET(COALESCE(receipts, '{}'), '$.\"%v\"', '%s') WHERE command_id = %v",
		p.cmd.currNonce,
		string(b),
		p.cmd.Tx.Id,
	)
	log.Debugf(query)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = p.db.ExecContext(ctx, query)
	return err
}

func (p *CmdProcessor) CallWithBoxTransactor(
	ctx context.Context,
	box *contract_gen.FuxPayBox,
	fn func(*contract_gen.FuxPayBoxTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithBoxTransactor(ctx, box, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithBoxFactoryTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxFactoryTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithBoxFactoryTransactor(ctx, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) WaitMined(ctx context.Context, tx *ethTypes.Transaction) (*ethTypes.Receipt, error) {
	nonce := p.cmd.currNonce
	if r, ok := p.cmd.receipts[string(nonce)]; ok {
		return r, nil
	}

	receipt, err := bind.WaitMined(ctx, p.fxClient.EthClient, tx)
	if err != nil {
		return nil, err
	}

	p.cmd.receipts[string(nonce)] = receipt
	log.Infof("update receipt to db: %s", receipt)
	p.updateReceipt(receipt)

	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return receipt, ErrTxExecuteFailed
	}

	return receipt, nil
}

// -----------------

type Executor interface {
	Execute(cmd Command) error
}

type EthExecutor struct {
	ethUrl        string
	contractAddrs g.ContractAddrs
	keystore      *keychain.Store
	db            *sql.DB
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

	var err error
	switch cmd.Tx.TxType {
	case SplitFX:
		err = e.splitFX(p)
	case Discount:
		err = e.pay(p)
	case Payment:
		err = e.pay(p)
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

	var err error
	switch cmd.Tx.TxType {
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := p.CallWithFxSplitTransactor(ctx, func(session *contract_gen.FuxSplitTransactorSession) error {
		var innerErr error
		tx, innerErr = session.SplitFux(&tokenId, newTokenIds, amounts, states)
		return innerErr
	})
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
	var innerErr error
	for idx, t := range output {
		boxId := generateBoxId(cmd.Tx.TxId, idx)
		to := t.Owner
		toAcc, err := e.keystore.GetAccount(to)
		if err != nil {
			log.Errorf("get account of %v failed: %v", to, err)
			return err
		}
		log.Infof("create box: %v for %s", boxId, to)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithBoxFactoryTransactor(
			ctx,
			func(session *contract_gen.FuxPayBoxFactoryTransactorSession) error {
				tx, innerErr = session.CreatePayBox(new(big.Int).SetUint64(boxId), toAcc.Address)
				return innerErr
			},
		)
		cancel()
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
		box, err := contract_gen.NewFuxPayBox(r.ContractAddress, p.fxClient.EthClient)
		if err != nil {
			log.Errorf("Create new FuxPayBox contract failed: %v", err)
			return err
		}

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithBoxTransactor(
			ctx,
			box,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				tx, innerErr = session.TransferOwnership(acc.Address)
				return innerErr
			},
		)
		cancel()
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := p.fxClient.CallWithBoxFactoryCaller(
		ctx,
		func(session *contract_gen.FuxPayBoxFactoryCallerSession) error {
			var innerErr error
			boxAddr, innerErr = session.GetPayBoxAddress(new(big.Int).SetUint64(boxId))
			return innerErr
		},
	)
	if err != nil {
		return err
	}
	box, err := contract_gen.NewFuxPayBox(boxAddr, p.fxClient.EthClient)
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
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err := p.CallWithBoxTransactor(
			ctx,
			box,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				var innerErr error
				tx, innerErr = session.Transfer(toAcc.Address, &t.Id)
				return innerErr
			},
		)
		cancel()
		if err != nil {
			log.Errorf("call FuxPayBox.Transfer contract failed: %v", err)
			return err
		}
	}

	log.Infof("destroy box: %v", boxId)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = p.CallWithBoxFactoryTransactor(ctx,
		func(session *contract_gen.FuxPayBoxFactoryTransactorSession) error {
			var innerErr error
			tx, innerErr = session.CloseBox(boxAddr)
			return innerErr
		})
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

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithFxTokenTransactor(ctx,
			func(session *contract_gen.FuxTokenTransactorSession) error {
				var innerErr error
				tx, innerErr = session.Mint(
					acc.Address,
					&t.Id,
					new(big.Int).SetUint64(t.Amount),
					new(big.Int).SetInt64(t.ExpireTime),
					new(big.Int).SetUint64(uint64(t.State)),
					"",
				)
				return innerErr
			})
		cancel()
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
	tokenIds := make([]*big.Int, batchSize)

	packing := func() error {
		jobId := new(big.Int).SetUint64(boxId*10 + uint64((consumedNum-1)/batchSize))
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err := p.CallWithFxBatchTransactor(ctx,
			func(session *contract_gen.FuxBatchTransactorSession) error {
				var innerErr error
				tx, innerErr = session.AddJob(jobId, fromAcc.Address, boxAddr, tokenIds)
				return innerErr
			})
		cancel()
		if err != nil {
			log.Errorf("call FuxBatch.AddJob contract failed: %v", err)
			return err
		}

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithFxBatchTransactor(ctx,
			func(session *contract_gen.FuxBatchTransactorSession) error {
				var innerErr error
				tx, innerErr = session.RunJob(jobId)
				return innerErr
			})
		cancel()
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
			packing()

			size := len(tokens) - consumedNum
			if size > batchSize {
				size = batchSize
			}
			tokenIds = make([]*big.Int, size)
		}
	}

	if actualAmount%batchSize > 0 {
		packing()
	}

	return consumedNum, nil
}

func generateBoxId(txId uint64, index int) uint64 {
	return txId*10 + uint64(index)
}
