package keychain

import (
	"testing"

	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	//rawUrl     = "http://10.26.1.17:62361"
	//key        = "86007224917ef7efd9e0d2c9cf3c85fb48f95ba55cc84fa8e4cdc6a5477e58bc"
	//passphrase = "dianrong"

	key        = "a1532f1be58eb3ad43d02f4c5f837f62ded6322886c9348702798f08d6cb751b"
	rawUrl     = "http://127.0.0.1:8545"
	passphrase = "123456"
)

var dbConfig = g.DbConfig{
	Schema:   "keystore",
	Host:     "127.0.0.1",
	Port:     "3306",
	Username: "root",
	Password: "chou1103",
}

var account Account

func init() {
	privKey, _ := crypto.HexToECDSA(key)
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	account = Account{address, key, passphrase}
}

//func TestCreateAccount(t *testing.T) {
//	store, err := NewStore(account, rawUrl, dir)
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer store.Close()
//
//	acc, err := store.CreateAccount("cf")
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	t.Logf("key: %s, addr: %s", acc.Key, acc.Address.Hex())
//}

func TestCreateAndGetAccount(t *testing.T) {
	store, err := NewStore(account, rawUrl, dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	accSource, err := store.CreateAccount("cf")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("key: %s, addr: %s", accSource.Key, accSource.Address.Hex())

	cid := "001"
	if err := store.StoreAccount(cid, accSource); err != nil {
		t.Fatal(err)
	}

	acc, err := store.GetAccount(cid)
	if err != nil {
		t.Fatal(err)
	}

	if acc.Key != accSource.Key {
		t.Fatalf("Key content mismatch: have key %s, want %s", acc.Key, accSource.Key)
	}
}
