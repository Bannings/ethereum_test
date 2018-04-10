package blockchain

import (
	"context"

	"gitlab.chainedfinance.com/chaincore/keychain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/eddyzhou/log"
)

type FxClient struct {
	ec    *ethclient.Client
	auth  *bind.TransactOpts
	nonce uint64
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

	c := &FxClient{
		ec:    conn,
		auth:  auth,
		nonce: nonce,
	}

	return c, nil
}
