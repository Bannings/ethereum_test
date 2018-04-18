package fx

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrTxExecuteFailed = errors.New("transaction execute failed")
)

type CmdProcessor struct {
	fxClient *blockchain.FxClient
	cmd      Command
}

func (p *CmdProcessor) CallWithFxTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxTokenTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithFxTransactor(ctx, fn)
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
		// TODO: write to db
	}
}

func (p *CmdProcessor) CallWithBoxTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithBoxTransactor(ctx, fn)
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

func (p *CmdProcessor) CallWithBoxFactoryCaller(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxFactoryCallerSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithBoxFactoryCaller(ctx, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxCaller(
	ctx context.Context,
	fn func(*contract_gen.FuxTokenCallerSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.GetNonce() {
		err := p.fxClient.CallWithFxCaller(ctx, fn)
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
	if r, ok := p.cmd.receipts[nonce]; ok {
		return r, nil
	}

	receipt, err := bind.WaitMined(ctx, p.fxClient.EthClient, tx)
	if err == nil && receipt.Status == ethTypes.ReceiptStatusSuccessful {
		p.cmd.receipts[nonce] = receipt
		// TODO: write to db
	}

	return receipt, err
}

// -----------------

type Executor interface {
	Execute(cmd Command) error
}

type EthExecutor struct {
	ethUrl        string
	contractAddrs g.ContractAddrs
	keystore      *keychain.Store
}

func (e *EthExecutor) Execute(cmd Command) error {
	log.Infof("execute command: %+v", cmd)

	switch cmd.Tx.TxType {
	case SplitFX, Discount, Payment:
		return e.executeBySupplier(cmd)
	case MintFX, Confirm:
		return e.executeByPlatform(cmd)
	default:
		return fmt.Errorf("unknow command: %v", cmd)
	}
}

func (e *EthExecutor) executeBySupplier(cmd Command) error {
	companyID := cmd.Tx.Sponsor()
	c, err := clientCache.GetOrCreate(companyID)
	if err != nil {
		log.Errorf("create personal client failed: %v", err)
		return err
	}
	p := &CmdProcessor{fxClient: c, cmd: cmd}

	switch cmd.Tx.TxType {
	case SplitFX:
		return e.splitFX(p)
	case Discount:
		return e.pay(p)
	case Payment:
		return e.pay(p)
	default:
		return fmt.Errorf("err command: %v", cmd)
	}
}

func (e *EthExecutor) executeByPlatform(cmd Command) error {
	acc := e.keystore.GetAdminAccount()
	c, err := blockchain.NewPersonalClient(e.ethUrl, acc, e.contractAddrs)
	if err != nil {
		log.Errorf("create admin client failed: %v", err)
		return err
	}
	p := &CmdProcessor{fxClient: c, cmd: cmd}

	switch cmd.Tx.TxType {
	case MintFX:
		return e.mintFX(p)
	case Confirm:
		return e.confirm(p)
	default:
		return fmt.Errorf("err command: %v", cmd)
	}
}

func (e *EthExecutor) splitFX(p *CmdProcessor) error {
	cmd := p.cmd
	token := cmd.Tx.Input[0]
	tokenId := token.ID

	var tokens [2]Token
	copy(tokens[:], cmd.Tx.Output[:2])

	var newTokenIds [2]*big.Int
	var amounts [2]*big.Int
	var states [2]*big.Int
	for i, t := range tokens {
		newTokenIds[i] = &t.ID
		amounts[i] = new(big.Int).SetUint64(t.Amount)
		states[i] = new(big.Int).SetInt64(int64(t.State))
	}
	log.Infof("--- split fx: tokenId: %v, newTokenIds: %+v, amounts: %+v", tokenId.String(), newTokenIds, amounts)

	var tx *ethTypes.Transaction
	var innerErr error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := p.CallWithFxTransactor(ctx, func(session *contract_gen.FuxTokenTransactorSession) error {
		tx, innerErr = session.SplitFux(&tokenId, newTokenIds, amounts, states)
		return innerErr
	})
	if err != nil {
		log.Errorf("call FuxToken.SplitFx contract failed: %v", err)
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
		boxID := generateBoxId(cmd.Tx.TxId, idx)
		log.Infof("create box: %v", boxID)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err := p.CallWithBoxFactoryTransactor(ctx,
			func(session *contract_gen.FuxPayBoxFactoryTransactorSession) error {
				tx, innerErr = session.CreatePayBox(new(big.Int).SetUint64(boxID))
				return innerErr
			})
		cancel()
		if err != nil {
			log.Errorf("call FuxPayBoxFactory.CreatePayBox contract failed: %v", err)
			return err
		}

		r, err := p.WaitMined(context.Background(), tx)
		if err != nil {
			log.Errorf("create payBox(id: %v) failed: %v", boxID, err)
			return err
		}

		n, err := e.boxing(p, t.Amount, input, r.ContractAddress)
		if err != nil {
			log.Errorf("transfer to box(id: %v) failed: %v", boxID, err)
			return err
		}

		log.Infof("transfer box(%v) to %v", boxID, t.Owner)
		acc := e.keystore.GetAdminAccount()

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithBoxTransactor(ctx,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				tx, innerErr = session.TransferOwnership(acc.Address)
				return innerErr
			})
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

func (e *EthExecutor) confirm(p *CmdProcessor) error {
	cmd := p.cmd
	txId := cmd.Tx.TxId
	boxId := generateBoxId(txId, 0)

	var boxAddr common.Address
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	p.CallWithBoxFactoryCaller(
		ctx,
		func(session *contract_gen.FuxPayBoxFactoryCallerSession) error {
			boxAddr, err = session.GetPayBoxAddress(new(big.Int).SetUint64(boxId))
			return err
		})

	input := cmd.Tx.Input
	to := cmd.Tx.Output[0].Owner
	toAcc, err := e.keystore.GetAccount(to)
	if err != nil {
		log.Errorf("get account of %v failed: %v", to, err)
		return err
	}

	var tx *ethTypes.Transaction
	var innerErr error
	for _, t := range input {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithBoxTransactor(ctx,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				tx, innerErr = session.Transfer(toAcc.Address, &t.ID)
				return innerErr
			})
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
	var innerErr error
	for _, t := range output {
		log.Infof("mint fx: id: %v, owner: %v, amount: %v, state: %v", t.ID.String(), t.Owner, t.Amount, t.State)
		acc, err := e.keystore.GetAccount(t.Owner)
		if err != nil {
			log.Errorf("get account of %v failed: %v", t.Owner, err)
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		err = p.CallWithFxTransactor(ctx,
			func(session *contract_gen.FuxTokenTransactorSession) error {
				tx, innerErr = session.CreateFux(
					acc.Address,
					&t.ID, new(big.Int).SetUint64(t.Amount),
					new(big.Int).SetInt64(t.ExpireTime),
					new(big.Int).SetUint64(uint64(t.State)),
					"")
				return innerErr
			})
		cancel()
		if err != nil {
			log.Errorf("call FuxToken.CreateFux contract failed: %v", err)
			return err
		}
	}
	return nil

}

func (e *EthExecutor) boxing(p *CmdProcessor, targetAmount uint64, tokens []Token, boxAddr common.Address) (int, error) {
	var actualAmount uint64
	var consumedNum int
	var tx *ethTypes.Transaction
	var innerErr error
	for _, t := range tokens {
		if actualAmount < targetAmount {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			err := p.CallWithFxTransactor(ctx,
				func(session *contract_gen.FuxTokenTransactorSession) error {
					tx, innerErr = session.Transfer(boxAddr, &t.ID)
					return innerErr
				})
			cancel()
			if err != nil {
				log.Errorf("call FuxToken.Transfer contract failed: %v", err)
				return 0, err
			}

			actualAmount += t.Amount
			consumedNum += 1
		}
	}

	// TODO: WaitMined?
	return consumedNum, nil
}

func generateBoxId(txId uint64, index int) uint64 {
	return txId*10 + uint64(index)
}
