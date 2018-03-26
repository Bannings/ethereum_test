package blockchain

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const (
	key      = `{"address":"0be438acbade790d5efd46efe199b59eef8e0cb8","crypto":{"cipher":"aes-128-ctr","ciphertext":"4c9e41ad9f9cebe547487f7d6b93dbd4b3195074e7854d95903ae1b3161ec67a","cipherparams":{"iv":"1ca5419b2236f534f8f22c46782b5a08"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"4541c0f43a0fddf944575befa2b93fa87468510665cf205082f9e9a5f20d1564"},"mac":"81ea3adb1c5dcd31c1e232dd54176515c165ed5e3c2db7a1564768166098b2bb"},"id":"962c374a-5bb8-46f4-b1d6-cd36ed044a9f","version":3}`
	password = "dianrong"
	rawUrl   = "http://10.26.1.17:62361"

	nodeContractAddr  = "0x4d07fd58970b6254bb500ace0a1d15b9394513af"
	storeContractAddr = "0x6aae7364a12fc276f454f42d90c8bdac0cdfb7f5"
)

var (
	conf = &Config{
		key,
		password,
		rawUrl,
		nodeContractAddr,
		storeContractAddr,
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
