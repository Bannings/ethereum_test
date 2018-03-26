package blockchain

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Key               string `json:"key"`
	Password          string `json:"password"`
	RawUrl            string `json:"rawUrl"`
	NodeContractAddr  string `json:"nodeContractAddr"`
	StoreContractAddr string `json:"storeContractAddr"`
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
