package fx

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	stateProcessing = "processing"
	stateProcessed  = "processed"
)

type CmdProcessor struct {
	fxClient *blockchain.FxClient
	cmd      Command
	db       *sql.DB
}

func (p *CmdProcessor) CallWithFxTokenTransactor(fn func(*contract_gen.FuxTokenTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		err := p.fxClient.CallWithFxTokenTransactor(fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxSplitTransactor(fn func(*contract_gen.FuxSplitTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		err := p.fxClient.CallWithFxSplitTransactor(fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxBatchTransactor(fn func(*contract_gen.FuxBatchTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		err := p.fxClient.CallWithFxBatchTransactor(fn)
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
		currNonce := p.fxClient.Nonce()
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

func (p *CmdProcessor) CallWithBoxTransactor(box *contract_gen.FuxPayBox, fn func(*contract_gen.FuxPayBoxTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		err := p.fxClient.CallWithBoxTransactor(box, fn)
		if err == nil {
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithBoxFactoryTransactor(fn func(*contract_gen.FuxPayBoxFactoryTransactorSession) error) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		err := p.fxClient.CallWithBoxFactoryTransactor(fn)
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

	receipt, err := bind.WaitMined(ctx, p.fxClient.EthClient(), tx)
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
