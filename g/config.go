package g

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"

	"gitlab.chainedfinance.com/chaincore/keychain"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	conf       *Config
	configLock = new(sync.RWMutex)
)

type Config struct {
	BlockchainConfig BlockChainConfig `json:"blockchain"`
	DbConfig         DbConfig         `json:"db"`
}

type DbConfig struct {
}

type BlockChainConfig struct {
	Key    string `json:"key"`
	RawUrl string `json:"rawUrl"`

	AdminKey        string `json:"adminKey"`
	AdminPassphrase string `json:"adminPassphrase"`

	ContractAddrs ContractAddrs `json:"contractAddrs"`
}

func (c *BlockChainConfig) ObtainAdminAccount() (keychain.Account, error) {
	configLock.RLock()
	defer configLock.RUnlock()

	if c.AdminKey == "" {
		return keychain.Account{}, errors.New("no admin account config")
	}

	key, err := crypto.HexToECDSA(c.AdminKey)
	if err != nil {
		return keychain.Account{}, err
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	return keychain.Account{Address: address, Key: c.AdminKey, Passphrase: c.AdminPassphrase}, nil
}

type ContractAddrs struct {
	FxTokenAddr      string `json:"fxTokenAddr"`
	FxPayBoxAddr     string `json:"fxPayBoxAddr"`
	FxBoxFactoryAddr string `json:"fxBoxFactoryAddr"`

	NodeAddr  string `json:"nodeAddr"`
	StoreAddr string `json:"storeAddr"`
}

func LoadConfig(jsonFile string) (*Config, error) {
	file, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	tmp := new(Config)
	if err = json.Unmarshal(file, tmp); err != nil {
		return nil, err
	}

	configLock.Lock()
	defer configLock.Unlock()
	conf = tmp
	return conf, nil
}

func GetConfig() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return conf
}
