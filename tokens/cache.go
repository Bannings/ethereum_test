package tokens

import (
	"sync"

	"gitlab.zhuronglink.com/chaincore/r2/blockchain"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"gitlab.zhuronglink.com/chaincore/r2/keychain"
	"gitlab.zhuronglink.com/infra/gocommons/lru"
)

var (
	defaultCache *ClientCache
	cacheOnce    sync.Once
)

func DefaultCache() *ClientCache {
	cacheOnce.Do(func() {
		conf := g.GetConfig()
		bConf := conf.BlockchainConfig
		cache := NewCache(bConf.RawUrl, keychain.DefaultStore(), 30)
		defaultCache = cache
	})
	return defaultCache
}

// ClientCache is a wrapper around an *lru.Cache that adds synchronization
// makes values always be FxClient
type ClientCache struct {
	rawUrl   string
	keystore *keychain.Store

	mu  sync.RWMutex
	lru *lru.Cache
}

func (c *ClientCache) FlushAll() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for c.lru.Len() > 0 {
		c.lru.RemoveOldest()
	}
}

func NewCache(rawUrl string, keystore *keychain.Store, maxEntries int) *ClientCache {
	c := new(ClientCache)
	c.lru = lru.New(maxEntries)
	c.lru.OnEvicted = func(key lru.Key, value interface{}) {
		entry := value.(*blockchain.TokenClient)
		c.handleEvicted(entry)
	}
	c.rawUrl = rawUrl
	c.keystore = keystore
	return c
}

func (c *ClientCache) handleEvicted(entry *blockchain.TokenClient) {
	entry.Close()
}

func (c *ClientCache) GetOrCreate(companyId string) (*blockchain.TokenClient, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var entry *blockchain.TokenClient
	val, ok := c.lru.Get(companyId)
	if !ok {
		acc, err := c.keystore.GetAccount(companyId)
		if err != nil {
			return nil, err
		}

		conf := g.GetConfig()
		entry, err = blockchain.NewPersonalClient(c.rawUrl, acc, conf.BlockchainConfig.ContractAddrs)
		if err != nil {
			return nil, err
		}

		c.lru.Put(companyId, entry)
	} else {
		entry = val.(*blockchain.TokenClient)
	}

	return entry, nil
}
