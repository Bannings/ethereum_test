package blockchain

import (
	"context"
	"math/big"
	"time"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gasLimit uint64 = 4712357
)

type FxClient struct {
	cli  *keychain.AccountClient
	auth *bind.TransactOpts

	fxToken      *contract_gen.FuxToken
	fxBoxFactory *contract_gen.FuxPayBoxFactory
	fxSpliter    *contract_gen.FuxSpliter
	fxBatch      *contract_gen.FuxBatch
	erc721Token  *contract_gen.ERC721Token
	fxLocker     *contract_gen.FuxLocker
	fxStorage    *contract_gen.FuxStorage
}

func NewTokenClient(cli *keychain.AccountClient, acccount keychain.Account, contractAddrs g.ContractAddrs) (*FxClient, error) {
	key, err := acccount.GetKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)

	fxToken, err := contract_gen.NewFuxToken(common.HexToAddress(contractAddrs.FxTokenAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxToken contract: %v", err)
		return nil, err
	}

	fxBoxFactory, err := contract_gen.NewFuxPayBoxFactory(common.HexToAddress(contractAddrs.FxBoxFactoryAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxBoxFactory contract: %v", err)
		return nil, err
	}

	fxSpliter, err := contract_gen.NewFuxSpliter(common.HexToAddress(contractAddrs.FxSpliterAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxSplit contract: %v", err)
		return nil, err
	}

	fxBatch, err := contract_gen.NewFuxBatch(common.HexToAddress(contractAddrs.FxBatchAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxSplit contract: %v", err)
		return nil, err
	}
	erc721Token, err := contract_gen.NewERC721Token(common.HexToAddress(contractAddrs.ERC721TokenAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a erc27Token contract: %v", err)
		return nil, err
	}
	fxLocker, err := contract_gen.NewFuxLocker(common.HexToAddress(contractAddrs.FXLockerAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxSplit contract: %v", err)
		return nil, err
	}
	fxStorage, err := contract_gen.NewFuxStorage(common.HexToAddress(contractAddrs.FXStorageAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a fxSplit contract: %v", err)
		return nil, err
	}
	c := &FxClient{
		cli:          cli,
		auth:         auth,
		fxToken:      fxToken,
		fxBoxFactory: fxBoxFactory,
		fxSpliter:    fxSpliter,
		fxBatch:      fxBatch,
		erc721Token:  erc721Token,
		fxLocker:     fxLocker,
		fxStorage:    fxStorage,
	}

	return c, nil
}

func NewPersonalClient(rawUrl string, acccount keychain.Account, contractAddrs g.ContractAddrs) (*FxClient, error) {
	cli, err := keychain.NewAccountClient(acccount, rawUrl, 5*time.Second)
	if err != nil {
		log.Errorf("new client failed: %v", err)
		return nil, err
	}

	c, err := NewTokenClient(cli, acccount, contractAddrs)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *FxClient) CallWithFxTokenTransactor(
	fn func(*contract_gen.FuxTokenTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxTokenTransactorSession{
		Contract: &c.fxToken.FuxTokenTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.cli.Nonce())),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	tx, err := fn(s)
	if err != nil {
		return nil, err
	}

	c.cli.IncrNonce()
	return tx, nil
}

func (c *FxClient) CallWithFxSpliterTransactor(
	fn func(*contract_gen.FuxSpliterTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxSpliterTransactorSession{
		Contract: &c.fxSpliter.FuxSpliterTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.cli.Nonce())),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	tx, err := fn(s)
	if err != nil {
		return nil, err
	}

	c.cli.IncrNonce()
	return tx, nil
}

func (c *FxClient) CallWithFxBatchTransactor(
	fn func(*contract_gen.FuxBatchTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxBatchTransactorSession{
		Contract: &c.fxBatch.FuxBatchTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.cli.Nonce())),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	tx, err := fn(s)
	if err != nil {
		return nil, err
	}

	c.cli.IncrNonce()
	return tx, nil
}

func (c *FxClient) CallWithBoxTransactor(
	box *contract_gen.FuxPayBox,
	fn func(*contract_gen.FuxPayBoxTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxPayBoxTransactorSession{
		Contract: &box.FuxPayBoxTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.cli.Nonce())),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	tx, err := fn(s)
	if err != nil {
		return nil, err
	}

	c.cli.IncrNonce()
	return tx, nil
}

func (c *FxClient) CallWithBoxFactoryTransactor(
	fn func(*contract_gen.FuxPayBoxFactoryTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxPayBoxFactoryTransactorSession{
		Contract: &c.fxBoxFactory.FuxPayBoxFactoryTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.cli.Nonce())),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	tx, err := fn(s)
	if err != nil {
		return nil, err
	}

	c.cli.IncrNonce()
	return tx, nil
}

func (c *FxClient) CallWithBoxFactoryCaller(fn func(*contract_gen.FuxPayBoxFactoryCallerSession) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.FuxPayBoxFactoryCallerSession{
		Contract: &c.fxBoxFactory.FuxPayBoxFactoryCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}

	return fn(s)
}

func (c *FxClient) CallWithFxTokenCaller(ctx context.Context, fn func(*contract_gen.FuxTokenCallerSession) error) error {
	s := &contract_gen.FuxTokenCallerSession{
		Contract: &c.fxToken.FuxTokenCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}

	return fn(s)
}

func (c *FxClient) CallWithERC721TokenCaller(ctx context.Context) *contract_gen.ERC721TokenCallerSession {
	session := &contract_gen.ERC721TokenCallerSession{
		Contract: &c.erc721Token.ERC721TokenCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
	return session
}

func (c *FxClient) CallWithFXLockerCaller(ctx context.Context) *contract_gen.FuxLockerCallerSession {
	session := &contract_gen.FuxLockerCallerSession{
		Contract: &c.fxLocker.FuxLockerCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
	return session
}

func (c *FxClient) CallWithFXStorageCaller(ctx context.Context) *contract_gen.FuxStorageCallerSession {
	session := &contract_gen.FuxStorageCallerSession{
		Contract: &c.fxStorage.FuxStorageCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
	return session
}

func (c *FxClient) Nonce() uint64 {
	return c.cli.Nonce()
}

func (c *FxClient) RefreshNonce() error {
	return c.cli.RefreshNonce()
}

func (c *FxClient) EthClient() *ethclient.Client {
	return c.cli.EthClient
}

func (c *FxClient) Close() {
	c.cli.Close()
}