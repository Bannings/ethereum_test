package fx

import (
	"fmt"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"

	"github.com/eddyzhou/log"
)

type Executor interface {
	Execute(cmd Command) error
}

type EthExecutor struct {
	cfClient blockchain.FxClient
}

func (e *EthExecutor) Execute(cmd Command) error {
	switch cmd.Type {
	case SplitFX:
		return e.splitFX(cmd)
	case MintFX:
		return e.mintFX(cmd)
	case Discount:
		return e.pay(cmd)
	case Payment:
		return e.pay(cmd)
	default:
		return fmt.Errorf("Unknow command: %v", cmd)
	}
}

func (e *EthExecutor) splitFX(cmd Command) error {
	token := cmd.T.Input[0]
	tokenId := token.ID
	var newTokenIds []uint64
	var amounts []uint64
	for _, t := range cmd.T.Output {
		newTokenIds = append(newTokenIds, t.ID)
		amounts = append(amounts, t.Amount)
	}

	log.Infof("--- split fx: tokenId: %v, newTokenIds: %+v, amounts: %+v", tokenId, newTokenIds, amounts)

	// TODO: call split contract

	return nil
}

func (e *EthExecutor) pay(cmd Command) error {
	output := cmd.T.Output
	input := cmd.T.Input
	for idx, t := range output {
		boxID := cmd.T.TxId*10 + uint64(idx)
		log.Infof("create box: %v", boxID)
		// TODO: create box

		n, err := e.consolidate(t.Amount, input, boxID)
		if err != nil {
			return err
		}

		to := t.Owner
		log.Infof("transfer box(%v) to %v", boxID, to)

		// TODO: transfer box to $to

		input = input[n:]
	}

	return nil
}

func (e *EthExecutor) mintFX(cmd Command) error {
	output := cmd.T.Output
	for _, t := range output {
		log.Infof("mint fx: id: %v, owner: %v, amount: %v, state: %v", t.ID, t.Owner, t.Amount, t.State)
		// TODO: mint fx
	}
	return nil

}

func (e *EthExecutor) consolidate(targetAmount uint64, tokens []Token, boxID uint64) (consumedNum int, err error) {
	var actualAmount uint64
	for _, t := range tokens {
		if actualAmount < targetAmount {
			// TODO: transfer token to box
			actualAmount += t.Amount
			consumedNum += 1
		}
	}
	return consumedNum, nil
}
