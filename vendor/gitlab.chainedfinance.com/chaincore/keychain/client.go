package keychain

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return big.NewInt(1000000000000000000)
}

type Client struct {
	account   Account
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	timeout   time.Duration
	nonce     uint64
}

func newClient(account Account, rawUrl string, timeout time.Duration) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c, err := rpc.DialContext(ctx, rawUrl)
	if err != nil {
		return nil, err
	}

	ethclient := ethclient.NewClient(c)

	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()
	nonce, err := ethclient.PendingNonceAt(ctx, account.Address)
	if err != nil {
		return nil, err
	}

	return &Client{account, c, ethclient, timeout, nonce}, nil
}

func (c *Client) PersonalImportRawKey(keyHex, passphrase string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	var addr string
	if err := c.rpcClient.CallContext(ctx, &addr, "personal_importRawKey", keyHex, passphrase); err != nil {
		return "", err
	}
	return addr, nil
}

func (c *Client) PersonalUnlockAccount(address string, passphrase string, duration int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	var succ bool
	if err := c.rpcClient.CallContext(ctx, &succ, "personal_unlockAccount", address, passphrase, duration); err != nil {
		return false, err
	}
	return succ, nil
}

func (c *Client) SendTransaction(tx *types.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.ethClient.SendTransaction(ctx, tx)
}

func (c *Client) PendingNonceAt(account common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.ethClient.PendingNonceAt(ctx, account)
}

func (c *Client) Close() {
	c.rpcClient.Close()
}
