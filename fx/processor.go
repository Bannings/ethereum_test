package fx

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"strconv"
)

const (
	stateProcessing = "processing"
	stateProcessed  = "processed"
)

var (
	ErrNoTxHash = errors.New("transation executed but not save txHash")
)

type CmdProcessor struct {
	fxClient *blockchain.FxClient
	cmd      Command
	db       *sql.DB
}

func (p *CmdProcessor) CallWithFxTokenTransactor(
	fn func(*contract_gen.FuxTokenTransactorSession) (*ethTypes.Transaction, error),
) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		tx, err := p.fxClient.CallWithFxTokenTransactor(fn)
		if err == nil {
			txHash := tx.Hash().Hex()
			p.cmd.txHashes[strconv.FormatUint(p.cmd.currNonce, 10)] = txHash
			p.saveTxHash(txHash)
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxSplitTransactor(
	fn func(*contract_gen.FuxSplitTransactorSession) (*ethTypes.Transaction, error),
) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		tx, err := p.fxClient.CallWithFxSplitTransactor(fn)
		if err == nil {
			txHash := tx.Hash().Hex()
			p.cmd.txHashes[strconv.FormatUint(p.cmd.currNonce, 10)] = txHash
			p.saveTxHash(txHash)
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithFxBatchTransactor(
	fn func(*contract_gen.FuxBatchTransactorSession) (*ethTypes.Transaction, error),
) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		tx, err := p.fxClient.CallWithFxBatchTransactor(fn)
		if err == nil {
			txHash := tx.Hash().Hex()
			p.cmd.txHashes[strconv.FormatUint(p.cmd.currNonce, 10)] = txHash
			p.saveTxHash(txHash)
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
			p.cmd.Tx.TxId,
			p.cmd.startNonce,
		)
		p.createProcedure()
	}
}

func (p *CmdProcessor) createProcedure() error {

	stmt, err := p.db.Prepare("INSERT INTO cmd_procedure(transaction_id, start_nonce, state) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, p.cmd.Tx.Id, p.cmd.startNonce, stateProcessing)
	if err != nil {
		return err
	}
	return nil
}

func (p *CmdProcessor) finishProcedure() error {
	b, err := json.Marshal(p.cmd.txHashes)
	if err != nil {
		return err
	}

	stmt, err := p.db.Prepare("UPDATE cmd_procedure SET tx_hashes = ?,state = ? WHERE transaction_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, string(b), stateProcessed, p.cmd.Tx.Id)
	return nil
}

func (p *CmdProcessor) saveTxHash(txHash string) error {
	query := fmt.Sprintf(
		"UPDATE cmd_procedure set tx_hashes = JSON_SET(COALESCE(tx_hashes, '{}'), '$.\"%v\"', '%s') WHERE transaction_id = %v",
		p.cmd.currNonce,
		txHash,
		p.cmd.Tx.TxId,
	)
	log.Debugf(query)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := p.db.ExecContext(ctx, query)
	return err
}

func (p *CmdProcessor) updateStatus() error {
	conn, err := p.db.Begin()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	stmt, err := p.db.Prepare("UPDATE transactions SET status = ? WHERE transaction_id = ?")
	if err != nil {
		conn.Rollback()
		return err
	}
	_, err = stmt.ExecContext(ctx, "failed", p.cmd.Tx.Id)

	if err != nil {
		conn.Rollback()
		return err
	}

	conn.Commit()
	return nil

}

func (p *CmdProcessor) CallWithBoxTransactor(
	box *contract_gen.FuxPayBox,
	fn func(*contract_gen.FuxPayBoxTransactorSession) (*ethTypes.Transaction, error),
) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		tx, err := p.fxClient.CallWithBoxTransactor(box, fn)
		if err == nil {
			txHash := tx.Hash().Hex()
			p.cmd.txHashes[string(p.cmd.currNonce)] = txHash
			p.saveTxHash(txHash)
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) CallWithBoxFactoryTransactor(
	fn func(*contract_gen.FuxPayBoxFactoryTransactorSession) (*ethTypes.Transaction, error),
) error {
	p.initCmdNonce()
	if p.cmd.currNonce >= p.fxClient.Nonce() {
		tx, err := p.fxClient.CallWithBoxFactoryTransactor(fn)
		if err == nil {
			txHash := tx.Hash().Hex()
			p.cmd.txHashes[string(p.cmd.currNonce)] = txHash
			p.saveTxHash(txHash)
			p.cmd.currNonce += 1
		}
		return err
	}
	p.cmd.currNonce += 1
	return nil
}

func (p *CmdProcessor) WaitMined(ctx context.Context, tx *ethTypes.Transaction) (*ethTypes.Receipt, error) {
	nonce := p.cmd.currNonce
	if txHash, ok := p.cmd.txHashes[string(nonce)]; ok {
		log.Infof("txhash:%v", txHash)
		return waitMined(ctx, p.fxClient.EthClient(), txHash)
	}

	if tx == nil {
		return nil, ErrNoTxHash
	}

	receipt, err := bind.WaitMined(ctx, p.fxClient.EthClient(), tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == ethTypes.ReceiptStatusFailed {
		return receipt, ErrTxExecuteFailed
	}

	return receipt, nil
}

func waitMined(ctx context.Context, b bind.DeployBackend, txHash string) (*ethTypes.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, _ := b.TransactionReceipt(ctx, common.HexToHash(txHash))
		if receipt != nil {
			return receipt, nil
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
