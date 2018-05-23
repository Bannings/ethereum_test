package fx

import (
	"database/sql"

	"github.com/ethereum/go-ethereum/crypto"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
)

var (
	db          *sql.DB
	bConf       g.BlockChainConfig
	bcConfig    g.BlockChainConfig
	keystore    *keychain.Store
	clientCache *ClientCache
	adminClient *blockchain.FxClient
)

func Init() {
	conf := g.GetConfig()
	bConf = conf.BlockchainConfig

	cfKey := conf.BlockchainConfig.AdminKey
	privKey, _ := crypto.HexToECDSA(cfKey)
	address := crypto.PubkeyToAddress(privKey.PublicKey)
	cfAccount := keychain.Account{Address: address, Key: cfKey}

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
	adminClient, err = blockchain.NewFxClient(cli, acc, bConf.ContractAddrs)
	if err != nil {
		panic(err)
	}
}
