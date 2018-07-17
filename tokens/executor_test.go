package tokens

import (
	"database/sql"
	"fmt"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	passphrase = "123456"
)

var (
	expireTime = time.Now().UnixNano()/1000000 + 365*24*3600*1000

	db          *sql.DB
	bConf       g.BlockChainConfig
	keystore    *keychain.Store
	clientCache *ClientCache
	adminClient *blockchain.FxClient
	txid        uint64
)

func init() {
	g.LoadConfig("../cf_dev.json")
	conf := g.GetConfig()
	bConf = conf.BlockchainConfig

	cfKey := conf.BlockchainConfig.AdminKey
	privKey, _ := crypto.HexToECDSA(cfKey)
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	cfAccount := keychain.Account{Address: address, Key: cfKey}
	txid = generateTxId()

	dbConfig := conf.DbConfig
	var err error
	if db, err = g.OpenDB(dbConfig); err != nil {
		panic(err)
	}

	ethUrl := conf.BlockchainConfig.RawUrl
	keystore, err = keychain.NewStore(cfAccount, ethUrl, dbConfig)
	if err != nil {
		panic(err)
	}

	clientCache = NewCache(bConf.RawUrl, keystore, 30)
	acc := keystore.GetAdminAccount()
	cli := keystore.GetAdminClient()
	adminClient, err = blockchain.NewTokenClient(cli, acc, bConf.ContractAddrs)
	if err != nil {
		panic(err)
	}
}

func TestMintFX(t *testing.T) {
	//_, err := settleIn(keystore, "supplier001")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//_, err = settleIn(keystore, "supplier002")
	//if err != nil {
	//	t.Fatal(err)
	//}

	tokenId := generateTokenId()
	inToken := Token{*big.NewInt(tokenId - 1), *big.NewInt(0), 14, "supplier001", "Normal", expireTime}
	token1 := Token{*big.NewInt(tokenId), *big.NewInt(tokenId - 1), 5, "supplier001", "Normal", expireTime}
	token2 := Token{*big.NewInt(tokenId + 1), *big.NewInt(tokenId - 1), 5, "supplier001", "Normal", expireTime}
	token3 := Token{*big.NewInt(tokenId + 2), *big.NewInt(tokenId - 1), 4, "supplier001", "Frozen", expireTime}
	tx := Transaction{Input: []Token{inToken}, Output: []Token{token1, token2, token3}, TxType: "MintFX", TxId: "", Id: 88}
	executor, err := NewEthExecutor(keystore, bConf, db, clientCache, adminClient)
	if err != nil {
		t.Fatal(err)
	}
	cmd := Command{Tx: tx, txHashes: make(map[string]string)}

	err = executor.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTransferFX(t *testing.T) {
	tokenId := generateTokenId()
	executor, err := NewEthExecutor(keystore, bConf, db, clientCache, adminClient)
	inToken1 := Token{*big.NewInt(tokenId), *big.NewInt(0), 140000, "supplier001", "Normal", expireTime}
	inToken2 := Token{*big.NewInt(tokenId + 1), *big.NewInt(0), 140000, "supplier001", "Normal", expireTime}
	token1 := Token{*big.NewInt(tokenId + 2), *big.NewInt(tokenId), 100000, "supplier001", "Normal", expireTime}
	token2 := Token{*big.NewInt(tokenId + 3), *big.NewInt(tokenId), 40000, "supplier001", "Frozen", expireTime}
	token3 := Token{*big.NewInt(tokenId + 4), *big.NewInt(tokenId + 1), 100000, "supplier001", "Normal", expireTime}
	token4 := Token{*big.NewInt(tokenId + 5), *big.NewInt(tokenId + 1), 40000, "supplier001", "Frozen", expireTime}
	tx1 := Transaction{Input: []Token{inToken1, inToken2}, Output: []Token{token1, token2, token3, token4}, TxType: "MintFX", TxId: "jgff", Id: 88}

	if err != nil {
		t.Fatal(err)
	}

	cmd := Command{Tx: tx1, txHashes: make(map[string]string)}

	err = executor.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}

	token5 := Token{*big.NewInt(tokenId + 6), *big.NewInt(tokenId + 2), 100000, "supplier002", "Normal", expireTime}
	token6 := Token{*big.NewInt(tokenId + 7), *big.NewInt(tokenId + 4), 60000, "supplier002", "Normal", expireTime}
	token7 := Token{*big.NewInt(tokenId + 8), *big.NewInt(tokenId + 4), 20000, "supplier002", "Normal", expireTime}
	token8 := Token{*big.NewInt(tokenId + 9), *big.NewInt(tokenId + 4), 20000, "supplier001", "Normal", expireTime}
	tx2 := Transaction{Input: []Token{token1, token3}, Output: []Token{token5, token6, token7, token8}, TxType: "Payment", TxId: "hkh", Id: 858}
	cmd = Command{Tx: tx2, txHashes: make(map[string]string)}
	err = executor.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBatchTransfer(t *testing.T) {
	tokenId := generateTokenId()
	executor, err := NewEthExecutor(keystore, bConf, db, clientCache, adminClient)
	inToken1 := Token{*big.NewInt(tokenId), *big.NewInt(0), 140000, "supplier001", "Normal", expireTime}
	token1 := Token{*big.NewInt(tokenId + 2), *big.NewInt(tokenId), 100000, "supplier001", "Normal", expireTime}
	token2 := Token{*big.NewInt(tokenId + 3), *big.NewInt(tokenId), 40000, "supplier001", "Frozen", expireTime}
	tx1 := Transaction{Input: []Token{inToken1}, Output: []Token{token1, token2}, TxType: "MintFX", TxId: "jgff", Id: 88}

	if err != nil {
		t.Fatal(err)
	}

	cmd := Command{Tx: tx1, txHashes: make(map[string]string)}

	err = executor.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}

	token3 := Token{*big.NewInt(tokenId + 7), *big.NewInt(tokenId + 3), 20000, "supplier002", "Normal", expireTime}
	token4 := Token{*big.NewInt(tokenId + 8), *big.NewInt(tokenId + 3), 20000, "supplier002", "Normal", expireTime}
	tx2 := Transaction{Input: []Token{token2}, Output: []Token{token3, token4}, TxType: "Payment", TxId: "hkh", Id: 858}
	cmd = Command{Tx: tx2, txHashes: make(map[string]string)}
	err = executor.Execute(cmd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSplitFX(t *testing.T) {
	tokenId := generateTokenId()
	executor, err := NewEthExecutor(keystore, bConf, db, clientCache, adminClient)
	if err != nil {
		t.Fatal(err)
	}
	token1 := Token{*big.NewInt(tokenId), *big.NewInt(0), 100, "supplier001", "Normal", expireTime}
	tx1 := Transaction{Input: []Token{token1}, Output: []Token{}, TxType: "MintFX"}
	cmd1 := Command{Tx: tx1, txHashes: make(map[string]string)}
	err = executor.Execute(cmd1)
	if err != nil {
		t.Logf("mint failed")
		t.Fatal(err)
	}
	token2 := Token{*big.NewInt(tokenId + 1), *big.NewInt(tokenId), 50, "supplier001", "Normal", expireTime}
	token3 := Token{*big.NewInt(tokenId + 2), *big.NewInt(tokenId), 50, "supplier001", "Frozen", expireTime}
	tx2 := Transaction{Input: []Token{token1}, Output: []Token{token2, token3}, TxType: "SplitFX"}
	cmd2 := Command{Tx: tx2, txHashes: make(map[string]string)}
	err = executor.Execute(cmd2)
	if err != nil {
		t.Logf("first split failed")
		t.Fatal(err)
	}
	token4 := Token{*big.NewInt(tokenId + 3), *big.NewInt(tokenId + 2), 30, "supplier001", "Normal", expireTime}
	token5 := Token{*big.NewInt(0), *big.NewInt(tokenId + 2), 20, "supplier001", "Normal", expireTime}
	tx3 := Transaction{Input: []Token{token3}, Output: []Token{token4, token5}, TxType: "SplitFX"}
	cmd3 := Command{Tx: tx3, txHashes: make(map[string]string)}
	err = executor.Execute(cmd3)
	if err != nil {
		t.Logf("second split failed")
		t.Fatal(err)
	}
}

func generateTokenId() int64 {
	return time.Now().UnixNano()
}

func generateTxId() uint64 {
	return uint64(time.Now().Unix())
}

func settleIn(store *keychain.Store, supplierId string) (keychain.Account, error) {
	acc, err := store.CreateAccount(passphrase)
	if err != nil {
		fmt.Printf("create account failed: %v", err)
		return keychain.Account{}, err
	}

	if err := store.StoreAccount(supplierId, acc); err != nil {
		fmt.Printf("store account failed: %v", err)
		return keychain.Account{}, err
	}

	return acc, nil
}

//eth.getTransactionReceipt("")
