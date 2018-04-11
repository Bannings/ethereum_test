package blockchain

import (
	"context"
	"math/big"
	"sync/atomic"

	"gitlab.chainedfinance.com/chaincore/contract-gen"
	"gitlab.chainedfinance.com/chaincore/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	FxTokenAddr, FxPayBoxAddr, FxBoxFactoryAddr string
)

type FxClient struct {
	EthClient *ethclient.Client
	Auth      *bind.TransactOpts
	nonce     uint64

	fxToken      *contract_gen.FuxToken
	fxPayBox     *contract_gen.FuxPayBox
	fxBoxFactory *contract_gen.FuxPayBoxFactory
}

func NewPersonalClient(rawUrl string, acccount keychain.Account) (*FxClient, error) {
	key, err := acccount.GetKey()
	if err != nil {
		return nil, err
	}

	conn, err := ethclient.Dial(rawUrl)
	if err != nil {
		log.Errorf("Fail to connect to the Ethereum client: %v", err)
		return nil, err
	}

	auth := bind.NewKeyedTransactor(key)
	nonce, err := conn.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	log.Debugf("nonce: %v", nonce)

	fxToken, err := contract_gen.NewFuxToken(common.HexToAddress(FxTokenAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a fxToken contract: %v", err)
		return nil, err
	}

	fxPayBox, err := contract_gen.NewFuxPayBox(common.HexToAddress(FxPayBoxAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a fxPayBox contract: %v", err)
		return nil, err
	}

	fxBoxFactory, err := contract_gen.NewFuxPayBoxFactory(common.HexToAddress(FxBoxFactoryAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a fxBoxFactory contract: %v", err)
		return nil, err
	}

	c := &FxClient{
		EthClient:    conn,
		Auth:         auth,
		nonce:        nonce,
		fxToken:      fxToken,
		fxPayBox:     fxPayBox,
		fxBoxFactory: fxBoxFactory,
	}

	return c, nil
}

func (c *FxClient) CallWithFxTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxTokenTransactorSession) error) error {

	s := &contract_gen.FuxTokenTransactorSession{
		Contract: &c.fxToken.FuxTokenTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.Auth.From,
			Signer:   c.Auth.Signer,
			Nonce:    big.NewInt(int64(c.nonce)),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	if err := fn(s); err != nil {
		return err
	}

	atomic.AddUint64(&c.nonce, 1)
	return nil
}

func (c *FxClient) CallWithBoxTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxTransactorSession) error) error {

	s := &contract_gen.FuxPayBoxTransactorSession{
		Contract: &c.fxPayBox.FuxPayBoxTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.Auth.From,
			Signer:   c.Auth.Signer,
			Nonce:    big.NewInt(int64(c.nonce)),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	if err := fn(s); err != nil {
		return err
	}

	atomic.AddUint64(&c.nonce, 1)
	return nil
}

func (c *FxClient) CallWithBoxFactoryTransactor(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxFactoryTransactorSession) error) error {

	s := &contract_gen.FuxPayBoxFactoryTransactorSession{
		Contract: &c.fxBoxFactory.FuxPayBoxFactoryTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.Auth.From,
			Signer:   c.Auth.Signer,
			Nonce:    big.NewInt(int64(c.nonce)),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	if err := fn(s); err != nil {
		return err
	}

	atomic.AddUint64(&c.nonce, 1)
	return nil
}

func (c *FxClient) CallWithBoxFactoryCaller(
	ctx context.Context,
	fn func(*contract_gen.FuxPayBoxFactoryCallerSession) error) error {

	s := &contract_gen.FuxPayBoxFactoryCallerSession{
		Contract: &c.fxBoxFactory.FuxPayBoxFactoryCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.Auth.From,
			Context: ctx,
		},
	}

	return fn(s)
}
