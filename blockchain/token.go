package blockchain

import (
	"context"
	"math/big"
	"time"

	"gitlab.zhuronglink.com/chaincore/contract-gen"
	"gitlab.com/ethereum_test/g"
	"gitlab.com/ethereum_test/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gasLimit uint64 = 4712357
)

type TokenClient struct {
	cli  *keychain.AccountClient
	auth *bind.TransactOpts

	token        *contract_gen.ZrlToken
	tokenSpliter *contract_gen.ZrlSpliter
	tokenBatch   *contract_gen.ZrlBatch
	erc721Token  *contract_gen.ERC721Token
	tokenLocker  *contract_gen.ZrlLocker
	tokenStorage *contract_gen.ZrlStorage
}

func NewTokenClient(cli *keychain.AccountClient, acccount keychain.Account, contractAddrs g.ContractAddrs) (*TokenClient, error) {
	key, err := acccount.GetKey()
	if err != nil {
		return nil, err
	}
	auth := bind.NewKeyedTransactor(key)

	token, err := contract_gen.NewZrlToken(common.HexToAddress(contractAddrs.TokenAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a Token contract: %v", err)
		return nil, err
	}

	tokenSpliter, err := contract_gen.NewZrlSpliter(common.HexToAddress(contractAddrs.SpliterAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a Split contract: %v", err)
		return nil, err
	}

	tokenBatch, err := contract_gen.NewZrlBatch(common.HexToAddress(contractAddrs.BatchAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a Split contract: %v", err)
		return nil, err
	}
	erc721Token, err := contract_gen.NewERC721Token(common.HexToAddress(contractAddrs.ERC721TokenAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a erc27Token contract: %v", err)
		return nil, err
	}
	tokenLocker, err := contract_gen.NewZrlLocker(common.HexToAddress(contractAddrs.LockerAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a Split contract: %v", err)
		return nil, err
	}
	tokenStorage, err := contract_gen.NewZrlStorage(common.HexToAddress(contractAddrs.StorageAddr), cli.EthClient)
	if err != nil {
		log.Errorf("Failed to instantiate a Split contract: %v", err)
		return nil, err
	}
	c := &TokenClient{
		cli:          cli,
		auth:         auth,
		token:        token,
		tokenSpliter: tokenSpliter,
		tokenBatch:   tokenBatch,
		erc721Token:  erc721Token,
		tokenLocker:  tokenLocker,
		tokenStorage: tokenStorage,
	}

	return c, nil
}

func NewPersonalClient(rawUrl string, acccount keychain.Account, contractAddrs g.ContractAddrs) (*TokenClient, error) {
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

func (c *TokenClient) CallWithTokenTransactor(
	fn func(*contract_gen.ZrlTokenTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.ZrlTokenTransactorSession{
		Contract: &c.token.ZrlTokenTransactor,
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

func (c *TokenClient) CallWithSpliterTransactor(
	fn func(*contract_gen.ZrlSpliterTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.ZrlSpliterTransactorSession{
		Contract: &c.tokenSpliter.ZrlSpliterTransactor,
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

func (c *TokenClient) CallWithBatchTransactor(
	fn func(*contract_gen.ZrlBatchTransactorSession) (*types.Transaction, error),
) (*types.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.cli.Timeout)
	defer cancel()
	s := &contract_gen.ZrlBatchTransactorSession{
		Contract: &c.tokenBatch.ZrlBatchTransactor,
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

func (c *TokenClient) CallWithTokenCaller(ctx context.Context, fn func(*contract_gen.ZrlTokenCallerSession) error) error {
	s := &contract_gen.ZrlTokenCallerSession{
		Contract: &c.token.ZrlTokenCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}

	return fn(s)
}

func (c *TokenClient) CallWithERC721TokenCaller(ctx context.Context) *contract_gen.ERC721TokenCallerSession {
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

func (c *TokenClient) CallWithLockerCaller(ctx context.Context) *contract_gen.ZrlLockerCallerSession {
	session := &contract_gen.ZrlLockerCallerSession{
		Contract: &c.tokenLocker.ZrlLockerCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
	return session
}

func (c *TokenClient) CallWithStorageCaller(ctx context.Context) *contract_gen.ZrlStorageCallerSession {
	session := &contract_gen.ZrlStorageCallerSession{
		Contract: &c.tokenStorage.ZrlStorageCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
	return session
}

func (c *TokenClient) Nonce() uint64 {
	return c.cli.Nonce()
}

func (c *TokenClient) RefreshNonce() error {
	return c.cli.RefreshNonce()
}

func (c *TokenClient) EthClient() *ethclient.Client {
	return c.cli.EthClient
}

func (c *TokenClient) Close() {
	c.cli.Close()
}
