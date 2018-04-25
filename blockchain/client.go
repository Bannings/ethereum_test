package blockchain

import (
	"context"
	"math/big"
	"sync/atomic"

	"gitlab.chainedfinance.com/chaincore/dr-contracts"
	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	gasLimit uint64 = 4712357
)

var DefaultClient StoreClient

type StoreClient interface {
	StoreCallerSession(ctx context.Context) *dr_contracts.ARStoreCallerSession
	StoreTransactorSession(ctx context.Context) *dr_contracts.ARStoreTransactorSession
	NodesCallerSession(ctx context.Context) *dr_contracts.CFNodesCallerSession
	NodesTransactorSession(ctx context.Context) *dr_contracts.CFNodesTransactorSession
}

type EthStoreClient struct {
	ec    *ethclient.Client
	auth  *bind.TransactOpts
	nonce uint64

	store *dr_contracts.ARStore
	nodes *dr_contracts.CFNodes
}

func (c *EthStoreClient) StoreTransactorSession(ctx context.Context) *dr_contracts.ARStoreTransactorSession {
	s := &dr_contracts.ARStoreTransactorSession{
		Contract: &c.store.ARStoreTransactor,
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

func (c *EthStoreClient) StoreCallerSession(ctx context.Context) *dr_contracts.ARStoreCallerSession {
	return &dr_contracts.ARStoreCallerSession{
		Contract: &c.store.ARStoreCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
}

func (c *EthStoreClient) NodesCallerSession(ctx context.Context) *dr_contracts.CFNodesCallerSession {
	return &dr_contracts.CFNodesCallerSession{
		Contract: &c.nodes.CFNodesCaller,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    c.auth.From,
			Context: ctx,
		},
	}
}

func (c *EthStoreClient) NodesTransactorSession(ctx context.Context) *dr_contracts.CFNodesTransactorSession {
	s := &dr_contracts.CFNodesTransactorSession{
		Contract: &c.nodes.CFNodesTransactor,
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

func NewEthStoreClient(conf g.BlockChainConfig) (*EthStoreClient, error) {
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

	store, err := dr_contracts.NewARStore(common.HexToAddress(conf.ContractAddrs.StoreAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a ArStore contract: %v", err)
		return nil, err
	}

	nodes, err := dr_contracts.NewCFNodes(common.HexToAddress(conf.ContractAddrs.NodeAddr), conn)
	if err != nil {
		log.Errorf("Failed to instantiate a CfNodes contract: %v", err)
		return nil, err
	}

	nonce, err := conn.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}
	log.Debugf("nonce: %v", nonce)

	c := &EthStoreClient{
		ec:    conn,
		auth:  auth,
		nonce: nonce,
		store: store,
		nodes: nodes,
	}

	return c, nil
}
