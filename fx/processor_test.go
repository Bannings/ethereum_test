package fx

import (
	"database/sql"
	"fmt"
	"math/big"
	"testing"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	passphrase = "123456"

	frozenRate = 0.2
	feeRate    = 0.001
)

var (
	expireTime = time.Now().UnixNano()/1000000 + 365*24*3600*1000
)

var (
	cfAccount     keychain.Account
	db            *sql.DB
	dbConfig      g.DbConfig
	ethUrl        string
	contractAddrs g.ContractAddrs
	keystore      *keychain.Store
)

func init() {
	g.LoadConfig("../cf_dev.json")
	conf := g.GetConfig()

	cfKey := conf.BlockchainConfig.AdminKey
	privKey, _ := crypto.HexToECDSA(cfKey)
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	cfAccount = keychain.Account{Address: address, Key: cfKey}

	dbConfig = conf.DbConfig
	contractAddrs = conf.BlockchainConfig.ContractAddrs
	var err error
	if db, err = g.OpenDB(dbConfig); err != nil {
		panic(err)
	}

	ethUrl = conf.BlockchainConfig.RawUrl
	keystore, err = keychain.NewStore(cfAccount, ethUrl, dbConfig)
	if err != nil {
		panic(err)
	}
	Init(ethUrl, keystore)
}

func TestClosure(t *testing.T) {
	var consumedNum int
	tokenIds := make([]*big.Int, 3)
	packing := func() {
		jobId := new(big.Int).SetUint64(uint64(consumedNum / 3))
		t.Logf("jobId: %v", jobId)
		t.Logf("tokens: %+v", tokenIds)
	}

	var actualAmount uint64
	tokens := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for idx, t := range tokens {
		actualAmount += 1
		consumedNum += 1
		i := idx % 3
		tokenIds[i] = new(big.Int).SetInt64(t)
		if actualAmount%3 == 0 {
			packing()

			size := len(tokens) - consumedNum
			if size > 3 {
				size = 3
			}
			tokenIds = make([]*big.Int, size)
		}
	}

	if actualAmount%3 > 0 {
		packing()
	}
}

func TestDB(t *testing.T) {
	tx := Transaction{Id: 1}
	cmd := Command{Tx: tx, startNonce: 1, receipts: make(map[string]*ethTypes.Receipt)}
	p := &CmdProcessor{cmd: cmd, db: db}
	if err := p.createProcedure(); err != nil {
		t.Fatal(err)
	}

	r := &ethTypes.Receipt{Status: 0}
	if err := p.updateReceipt(r); err != nil {
		t.Fatal(err)
	}

	r.Status = 1
	p.cmd.currNonce += 1
	if err := p.updateReceipt(r); err != nil {
		t.Fatal(err)
	}

	if err := p.finishProcedure(); err != nil {
		t.Fatal(err)
	}
}

func TestFX(t *testing.T) {
	defer keystore.Close()

	_, err := settleIn(keystore, "supplier001")
	if err != nil {
		t.Fatal(err)
	}
	_, err = settleIn(keystore, "supplier002")
	if err != nil {
		t.Fatal(err)
	}

	executor := &EthExecutor{ethUrl: ethUrl, contractAddrs: contractAddrs, keystore: keystore, db: db}

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
	t := Transaction{Input: input, Output: output, TxId: txId, TxType: Confirm}
	cmd := Command{Tx: t}
	return executor.Execute(cmd)
}

func pay(executor *EthExecutor, inputIds [2]big.Int, txId uint64) error {
	token1 := Token{inputIds[0], 50000, "supplier001", Normal, expireTime}
	token2 := Token{inputIds[1], 50, "supplier001", Normal, expireTime}
	input := []Token{token1, token2}

	token3 := Token{inputIds[0], 50000, "cf", Normal, expireTime}
	token4 := Token{inputIds[1], 50, "cf", Normal, expireTime}
	output := []Token{token3, token4}

	t := Transaction{Input: input, Output: output, TxId: txId, TxType: Payment}
	cmd := Command{Tx: t}
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
	t := Transaction{Input: input, Output: output, TxId: 0, TxType: SplitFX}
	cmd := Command{Tx: t}

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
	t := Transaction{Input: input, Output: output, TxId: 0, TxType: SplitFX}
	cmd := Command{Tx: t}

	err := executor.Execute(cmd)
	return tokenId1, tokenId2, err
}

func mintFX(executor *EthExecutor, tokenId int64, amount uint64) error {
	frozenAmount := uint64(float64(amount) * frozenRate)
	activeAmount := amount - frozenAmount

	token1 := Token{*big.NewInt(tokenId), activeAmount, "supplier001", Normal, expireTime}
	token2 := Token{*big.NewInt(tokenId + 1), frozenAmount, "supplier001", Normal, expireTime}
	tokens := [2]Token{token1, token2}

	t := Transaction{Output: tokens[:], TxType: MintFX}
	cmd := Command{Tx: t}
	return executor.Execute(cmd)
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
