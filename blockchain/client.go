package blockchain

import (
	"context"
	"math/big"
	"sync/atomic"

	"gitlab.chainedfinance.com/chaincore/dr-contracts"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gasLimit uint64 = 30000000
)

var DefaultClient Client

type Client interface {
	StoreCallerSession(ctx context.Context) *contracts.ArStoreCallerSession
	StoreTransactorSession(ctx context.Context) *contracts.ArStoreTransactorSession
	NodesCallerSession(ctx context.Context) *contracts.CfNodesCallerSession
	NodesTransactorSession(ctx context.Context) *contracts.CfNodesTransactorSession
}

type EthClient struct {
	ec    *ethclient.Client
	auth  *bind.TransactOpts
	nonce uint64

	store *contracts.ArStore
	nodes *contracts.CfNodes
}

func (c *EthClient) StoreTransactorSession(ctx context.Context) *contracts.ArStoreTransactorSession {
	s := &contracts.ArStoreTransactorSession{
		Contract: &c.store.ArStoreTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.nonce)),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	atomic.AddUint64(&c.nonce, 1)
	return s
}

func (c *EthClient) StoreCallerSession(ctx context.Context) *contracts.ArStoreCallerSession {
	return &contracts.ArStoreCallerSession{
		Contract: &c.store.ArStoreCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
}

func (c *EthClient) NodesCallerSession(ctx context.Context) *contracts.CfNodesCallerSession {
	return &contracts.CfNodesCallerSession{
		Contract: &c.nodes.CfNodesCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
}

func (c *EthClient) NodesTransactorSession(ctx context.Context) *contracts.CfNodesTransactorSession {
	s := &contracts.CfNodesTransactorSession{
		Contract: &c.nodes.CfNodesTransactor,
		TransactOpts: bind.TransactOpts{
			From:     c.auth.From,
			Signer:   c.auth.Signer,
			Nonce:    big.NewInt(int64(c.nonce)),
			GasLimit: gasLimit,
			Context:  ctx,
		},
	}

	atomic.AddUint64(&c.nonce, 1)
	return s
}

func NewEthClient(conf *Config) (*EthClient, error) {
	privKey, err := crypto.HexToECDSA(conf.Key)
	if err != nil {
		log.Errorf("Failed to convert private key: %v", err)
		return nil, err
	}

	conn, err := ethclient.Dial(conf.RawUrl)
	if err != nil {
		log.Errorf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	store, err := contracts.NewArStore(common.HexToAddress(conf.ContractAddrs.StoreAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a ArStore contract: %v", err)
		return nil, err
	}

	nodes, err := contracts.NewCfNodes(common.HexToAddress(conf.ContractAddrs.NodeAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a CfNodes contract: %v", err)
		return nil, err
	}

	nonce, err := conn.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	log.Debugf("nonce: %v", nonce)

	c := &EthClient{
		ec:    conn,
		auth:  auth,
		nonce: nonce,
		store: store,
		nodes: nodes,
	}

	return c, nil
}
