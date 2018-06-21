package keychain

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type AccountClient struct {
	account   Account
	rpcClient *rpc.Client
	EthClient *ethclient.Client
	Timeout   time.Duration
	nonce     *uint64
}

func NewAccountClient(account Account, rawUrl string, timeout time.Duration) (*AccountClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c, err := rpc.DialContext(ctx, rawUrl)
	if err != nil {
		log.Errorf("dial rawUrl:%s failed: %s", rawUrl, err)
		return nil, err
	}

	ethClient := ethclient.NewClient(c)

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	nonce, err := ethClient.PendingNonceAt(ctx, account.Address)
	if err != nil {
		log.Errorf("get pending nonce failed: %v", err)
		return nil, err
	}

	return &AccountClient{account, c, ethClient, timeout, &nonce}, nil
}

func (c *AccountClient) PersonalImportRawKey(keyHex, passphrase string) (string, error) {
	key := keyHex
	if keyHex[0:2] != "0x" && keyHex[0:2] != "0X" {
		key = "0x" + keyHex
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	var addr string
	if err := c.rpcClient.CallContext(ctx, &addr, "personal_importRawKey", key, passphrase); err != nil {
		return "", err
	}
	return addr, nil
}

func (c *AccountClient) PersonalUnlockAccount(address string, passphrase string, duration int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	var succ bool
	if err := c.rpcClient.CallContext(ctx, &succ, "personal_unlockAccount", address, passphrase, duration); err != nil {
		return false, err
	}
	return succ, nil
}

func (c *AccountClient) SendTransaction(tx *types.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	return c.EthClient.SendTransaction(ctx, tx)
}

func (c *AccountClient) PendingNonceAt(account common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	return c.EthClient.PendingNonceAt(ctx, account)
}

func (c *AccountClient) Nonce() uint64 {
	return *c.nonce
}

func (c *AccountClient) IncrNonce() {
	atomic.AddUint64(c.nonce, 1)
}

func (c *AccountClient) RefreshNonce() error {
	nonce, err := c.PendingNonceAt(c.account.Address)
	if err != nil {
		return err
	}

	atomic.StoreUint64(c.nonce, nonce)
	return nil
}

func (c *AccountClient) Close() {
	c.rpcClient.Close()
}
