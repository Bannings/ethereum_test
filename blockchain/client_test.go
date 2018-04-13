package blockchain

import (
	"context"
	"testing"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	key    = `86007224917ef7efd9e0d2c9cf3c85fb48f95ba55cc84fa8e4cdc6a5477e58bc`
	rawUrl = "http://10.26.1.17:62361"

	nodeContractAddr  = "0x4d07fd58970b6254bb500ace0a1d15b9394513af"
	storeContractAddr = "0x6aae7364a12fc276f454f42d90c8bdac0cdfb7f5"
)

var (
	conf = &g.BlockChainConfig{
		key,
		rawUrl,
		"",
		"",
		g.ContractAddrs{NodeAddr: nodeContractAddr, StoreAddr: storeContractAddr},
	}
)

func init() {
	DefaultClient, _ = NewEthClient(conf)
}

func TestExistedNode(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pubPEM, info, err := DefaultClient.NodesCallerSession(ctx).GetNodeKey("cf")
	if err != nil {
		t.Errorf("Failed to get node key: %v", err)
	}
	t.Logf("node key: %s, %s", pubPEM, info)
}

func TestAddNode(t *testing.T) {
	tx, err := DefaultClient.NodesTransactorSession(context.Background()).DeleteNode("cf2")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := bind.WaitMined(context.Background(), DefaultClient.(*EthClient).ec, tx); err != nil {
		t.Fatal(err)
	}

	pubPEM := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDY6xEcCugrz4bni+8W56iqzsc1
BS2j5HT5M8u8w9O5YfCc8lBbeUL0alahzggtVWowHG2Y+NyWmPY94cLB3smYm+LP
UxQjC0Ku3H05PGft9Gmm2YkfLSnAAitSO5hLKqSE0OShVyc/RA02GLyV4f51uFCU
U5Al54VGuEfJQ5y8PQIDAQAB
-----END PUBLIC KEY-----`
	tx, err = DefaultClient.NodesTransactorSession(context.Background()).AddNode("cf2", pubPEM, "cf2 info")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := bind.WaitMined(context.Background(), DefaultClient.(*EthClient).ec, tx); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pubPEM, info, err := DefaultClient.NodesCallerSession(ctx).GetNodeKey("cf2")
	if err != nil {
		t.Errorf("Failed to get node key: %v", err)
	}
	t.Logf("node key: %s, %s", pubPEM, info)
}

func TestContract(t *testing.T) {
	contractNo := "test_001"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := DefaultClient.StoreCallerSession(ctx).CheckContractNo(contractNo)
	if !existed {
		t.Errorf("Check contractNo failed, got: %v, want: %v.", existed, true)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	contractNo, info, err := DefaultClient.StoreCallerSession(ctx).QueryContractInfo(contractNo)
	if err != nil {
		t.Errorf("Failed to query contract info: %v", err)
	}
	t.Logf("contract key: %s, %s", contractNo, info)
}
