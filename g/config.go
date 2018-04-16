package g

import (
	"encoding/json"
	"io/ioutil"
	"sync"
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
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BlockChainConfig struct {
	Key    string `json:"key"`
	RawUrl string `json:"rawUrl"`

	AdminKey        string `json:"adminKey"`
	AdminPassphrase string `json:"adminPassphrase"`

	ContractAddrs ContractAddrs `json:"contractAddrs"`
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
