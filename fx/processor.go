package fx

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/keychain"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type Executor interface {
	Execute(cmd Command) error
}

type EthExecutor struct {
	ethUrl   string
	keystore *keychain.Store
}

func (e *EthExecutor) Execute(cmd Command) error {
	log.Infof("execute command: %+v", cmd)

	switch cmd.Type {
	case SplitFX, Discount, Payment:
		return e.executeBySupplier(cmd)
	case MintFX, Confirm:
		return e.executeByPlatform(cmd)
	default:
		return fmt.Errorf("Unknow command: %v", cmd)
	}
}

func (e *EthExecutor) executeBySupplier(cmd Command) error {
	companyID := cmd.T.Sponsor()
	acc, err := e.keystore.GetAccount(companyID)
	if err != nil {
		log.Errorf("get account failed: %v", err)
		return err
	}

	c, err := blockchain.NewPersonalClient(e.ethUrl, acc)
	if err != nil {
		log.Errorf("create personal client failed: %v", err)
		return err
	}

	switch cmd.Type {
	case SplitFX:
		return e.splitFX(c, cmd)
	case Discount:
		return e.pay(c, cmd)
	case Payment:
		return e.pay(c, cmd)
	default:
		return fmt.Errorf("err command: %v", cmd)
	}
}

func (e *EthExecutor) executeByPlatform(cmd Command) error {
	acc := e.keystore.GetAdminAccount()
	c, err := blockchain.NewPersonalClient(e.ethUrl, acc)
	if err != nil {
		log.Errorf("create admin client failed: %v", err)
		return err
	}

	switch cmd.Type {
	case MintFX:
		return e.mintFX(c, cmd)
	case Confirm:
		return e.confirm(c, cmd)
	default:
		return fmt.Errorf("err command: %v", cmd)
	}
}

func (e *EthExecutor) splitFX(c *blockchain.FxClient, cmd Command) error {
	token := cmd.T.Input[0]
	tokenId := token.ID

	var tokens [2]Token
	copy(tokens[:], cmd.T.Output[:2])

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
	err := c.CallWithFxTransactor(ctx, func(session *contract_gen.FuxTokenTransactorSession) error {
		tx, innerErr = session.SplitFux(&tokenId, newTokenIds, amounts, states)
		return innerErr
	})
	if err != nil {
		log.Errorf("call FuxToken.SplitFx contract failed: %v", err)
		return err
	}

	_, err = bind.WaitMined(context.Background(), c.EthClient, tx)
	return err
}

func (e *EthExecutor) pay(c *blockchain.FxClient, cmd Command) error {
	output := cmd.T.Output
	input := cmd.T.Input

	var tx *ethTypes.Transaction
	var innerErr error
	for idx, t := range output {
		boxID := generateBoxId(cmd.T.TxId, idx)
		log.Infof("create box: %v", boxID)

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err := c.CallWithBoxFactoryTransactor(ctx,
			func(session *contract_gen.FuxPayBoxFactoryTransactorSession) error {
				tx, innerErr = session.CreatePayBox(new(big.Int).SetUint64(boxID))
				return innerErr
			})
		if err != nil {
			log.Errorf("call FuxPayBoxFactory.CreatePayBox contract failed: %v", err)
			return err
		}

		r, err := bind.WaitMined(context.Background(), c.EthClient, tx)
		if err != nil {
			log.Errorf("create payBox(id: %v) failed: %v", boxID, err)
			return err
		}

		n, err := e.boxing(c, t.Amount, input, r.ContractAddress)
		if err != nil {
			log.Errorf("transfer to box(id: %v) failed: %v", boxID, err)
			return err
		}

		log.Infof("transfer box(%v) to %v", boxID, t.Owner)
		acc := e.keystore.GetAdminAccount()

		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = c.CallWithBoxTransactor(ctx,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				tx, innerErr = session.TransferOwnership(acc.Address)
				return innerErr
			})
		if err != nil {
			log.Errorf("call FuxPayBox.TransferOwnership contract failed: %v", err)
			return err
		}

		input = input[n:]
	}

	_, err := bind.WaitMined(context.Background(), c.EthClient, tx)
	return err
}

func (e *EthExecutor) confirm(c *blockchain.FxClient, cmd Command) error {
	txId := cmd.T.TxId
	boxId := generateBoxId(txId, 0)

	var boxAddr common.Address
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	c.CallWithBoxFactoryCaller(
		ctx,
		func(session *contract_gen.FuxPayBoxFactoryCallerSession) error {
			boxAddr, err = session.GetPayBoxAddress(new(big.Int).SetUint64(boxId))
			return err
		})

	input := cmd.T.Input
	to := cmd.T.Output[0].Owner
	toAcc, err := e.keystore.GetAccount(to)
	if err != nil {
		log.Errorf("get account of %v failed: %v", to, err)
		return err
	}

	var tx *ethTypes.Transaction
	var innerErr error
	for _, t := range input {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = c.CallWithBoxTransactor(ctx,
			func(session *contract_gen.FuxPayBoxTransactorSession) error {
				tx, innerErr = session.Transfer(toAcc.Address, &t.ID)
				return innerErr
			})
		if err != nil {
			log.Errorf("call FuxPayBox.Transfer contract failed: %v", err)
			return err
		}
	}

	log.Infof("destroy box: %v", boxId)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = c.CallWithBoxFactoryTransactor(ctx,
		func(session *contract_gen.FuxPayBoxFactoryTransactorSession) error {
			tx, innerErr = session.CloseBox(boxAddr)
			return innerErr
		})
	if err != nil {
		log.Errorf("call FuxPayBoxFactory.CloseBox contract failed: %v", err)
		return err
	}

	_, err = bind.WaitMined(context.Background(), c.EthClient, tx)
	return err
}

func (e *EthExecutor) mintFX(c *blockchain.FxClient, cmd Command) error {
	output := cmd.T.Output
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
		defer cancel()
		err = c.CallWithFxTransactor(ctx,
			func(session *contract_gen.FuxTokenTransactorSession) error {
				tx, innerErr = session.CreateFux(
					acc.Address,
					&t.ID, new(big.Int).SetUint64(t.Amount),
					new(big.Int).SetInt64(t.ExpireTime),
					new(big.Int).SetUint64(uint64(t.State)),
					"")
				return innerErr
			})
		if err != nil {
			log.Errorf("call FuxToken.CreateFux contract failed: %v", err)
			return err
		}
	}
	return nil

}

func (e *EthExecutor) boxing(c *blockchain.FxClient, targetAmount uint64, tokens []Token, boxAddr common.Address) (int, error) {
	var actualAmount uint64
	var consumedNum int
	var tx *ethTypes.Transaction
	var innerErr error
	for _, t := range tokens {
		if actualAmount < targetAmount {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			err := c.CallWithFxTransactor(ctx,
				func(session *contract_gen.FuxTokenTransactorSession) error {
					tx, innerErr = session.Transfer(boxAddr, &t.ID)
					return innerErr
				})
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
