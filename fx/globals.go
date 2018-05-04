package fx

import (
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
)

var (
	clientCache *ClientCache
	adminClient *blockchain.FxClient
)

func Init(rawUrl string, keystore *keychain.Store) error {
	clientCache = NewCache(rawUrl, keystore, 30)

	acc := keystore.GetAdminAccount()
	cli := keystore.GetAdminClient()
	conf := g.GetConfig()
	var err error
	adminClient, err = blockchain.NewFxClient(cli, acc, conf.BlockchainConfig.ContractAddrs)
	return err
}

func Close() {
	clientCache.FlushAll()
	adminClient.Close()
}
