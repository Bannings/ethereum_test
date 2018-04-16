package fx

import (
	"math/big"
	"testing"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	cfKey      = "a1532f1be58eb3ad43d02f4c5f837f62ded6322886c9348702798f08d6cb751b"
	ethUrl     = "http://127.0.0.1:8545"
	passphrase = "123456"
)

var (
	expireTime = time.Now().UnixNano()/1000000 + 365*24*3600*1000
	frozenRate = 0.2
	feeRate    = 0.001

	dbConfig = g.DbConfig{
		Schema:   "keystore",
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: "chou1103",
	}
)

var cfAccount keychain.Account
var contractAddrs g.ContractAddrs

func init() {
	privKey, _ := crypto.HexToECDSA(cfKey)
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	cfAccount = keychain.Account{Address: address, Key: cfKey}

	contractAddrs = g.ContractAddrs{
		FxTokenAddr:      "0x77227767836175e4799262607Cbe2F9957fE5B0E",
		FxPayBoxAddr:     "0x5DE012D26771173038138c30764b4E62F5D643df",
		FxBoxFactoryAddr: "0xaFE3FfE684ff35436195D9947c7eEB002E40C60B",
	}
}

func TestFX(t *testing.T) {
	store, err := keychain.NewStore(cfAccount, ethUrl, dbConfig)
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	_, err = settleIn(store, "supplier001")
	if err != nil {
		t.Fatal(err)
	}
	_, err = settleIn(store, "supplier002")
	if err != nil {
		t.Fatal(err)
	}

	executor := &EthExecutor{ethUrl, contractAddrs, store}

	tokenId1 := generateTokenId()
	err = mintFX(executor, tokenId1, 100000)
	if err != nil {
		t.Fatal(err)
	}

	tokenId2, tokenId3, err := splitFX4Pay(executor, tokenId1)
	if err != nil {
		t.Fatal(err)
	}
	tokenId4, _, err := splitFX4Fee(executor, tokenId3)
	if err != nil {
		t.Fatal(err)
	}

	txId := time.Now().UnixNano()
	input := [2]big.Int{*tokenId2, *tokenId4}
	err = pay(executor, input, uint64(txId))
	if err != nil {
		t.Fatal(err)
	}

	err = confirm(executor, *tokenId2, uint64(txId))
	if err != nil {
		t.Fatal(err)
	}
}

func confirm(executor *EthExecutor, inputId big.Int, txId uint64) error {
	token1 := Token{inputId, 50000, "cf", Normal, expireTime}
	token2 := Token{inputId, 50000, "supplier002", Normal, expireTime}
	input := []Token{token1}
	output := []Token{token2}
	t := Transaction{input, output, txId}
	cmd := Command{t, Confirm}
	return executor.Execute(cmd)
}

func pay(executor *EthExecutor, inputIds [2]big.Int, txId uint64) error {
	token1 := Token{inputIds[0], 50000, "supplier001", Normal, expireTime}
	token2 := Token{inputIds[1], 50, "supplier001", Normal, expireTime}
	input := []Token{token1, token2}

	token3 := Token{inputIds[0], 50000, "cf", Normal, expireTime}
	token4 := Token{inputIds[1], 50, "cf", Normal, expireTime}
	output := []Token{token3, token4}

	t := Transaction{input, output, txId}
	cmd := Command{t, Payment}
	return executor.Execute(cmd)
}

func generateTokenId() int64 {
	return time.Now().UnixNano() / 1000000
}

func splitFX4Pay(executor *EthExecutor, inTokenID int64) (*big.Int, *big.Int, error) {
	inToken := Token{*big.NewInt(inTokenID), 100000 - uint64(100000*frozenRate), "supplier001", Normal, expireTime}

	tokenId1 := big.NewInt(generateTokenId())
	tokenId2 := big.NewInt(generateTokenId())
	outToken1 := Token{*tokenId1, 50000, "supplier001", Normal, expireTime}
	outToken2 := Token{*tokenId2, 30000, "supplier001", Normal, expireTime}

	input := []Token{inToken}
	output := []Token{outToken1, outToken2}
	t := Transaction{input, output, 0}
	cmd := Command{t, SplitFX}

	err := executor.Execute(cmd)
	return tokenId1, tokenId2, err
}

func splitFX4Fee(executor *EthExecutor, inTokenID *big.Int) (*big.Int, *big.Int, error) {
	fee := uint64(50000 * feeRate)
	inToken := Token{*inTokenID, 30000, "supplier001", Normal, expireTime}

	tokenId1 := big.NewInt(generateTokenId())
	tokenId2 := big.NewInt(generateTokenId())
	outToken1 := Token{*tokenId1, 30000 - fee, "supplier001", Normal, expireTime}
	outToken2 := Token{*tokenId2, fee, "supplier001", Normal, expireTime}

	input := []Token{inToken}
	output := []Token{outToken1, outToken2}
	t := Transaction{input, output, 0}
	cmd := Command{t, SplitFX}

	err := executor.Execute(cmd)
	return tokenId1, tokenId2, err
}

func mintFX(executor *EthExecutor, tokenId int64, amount uint64) error {
	frozenAmount := uint64(float64(amount) * frozenRate)
	activeAmount := amount - frozenAmount

	token1 := Token{*big.NewInt(tokenId), activeAmount, "supplier001", Normal, expireTime}
	token2 := Token{*big.NewInt(tokenId + 1), frozenAmount, "supplier001", Normal, expireTime}
	tokens := [2]Token{token1, token2}

	t := Transaction{Output: tokens[:]}
	cmd := Command{t, MintFX}
	return executor.Execute(cmd)
}

func settleIn(store *keychain.Store, supplierId string) (keychain.Account, error) {
	acc, err := store.CreateAccount(passphrase)
	if err != nil {
		log.Errorf("create account failed: %v", err)
		return keychain.Account{}, err
	}

	if err := store.StoreAccount(supplierId, acc); err != nil {
		log.Errorf("store account failed: %v", err)
		return keychain.Account{}, err
	}

	return acc, nil
}
