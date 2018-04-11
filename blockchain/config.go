package blockchain

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	"gitlab.chainedfinance.com/chaincore/keychain"

	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
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

	config := new(Config)
	if err = json.Unmarshal(file, config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) ObtainAdminAccount() (keychain.Account, error) {
	if c.AdminKey == "" {
		return keychain.Account{}, errors.New("No admin account config")
	}

	key, err := crypto.HexToECDSA(c.AdminKey)
	if err != nil {
		return keychain.Account{}, err
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	return keychain.Account{address, c.AdminKey, c.AdminPassphrase}, nil
}
