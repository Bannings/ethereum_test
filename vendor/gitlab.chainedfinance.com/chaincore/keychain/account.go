package keychain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	Address    common.Address `json:"address"`
	Key        string         `json:"key"`
	Passphrase string         `json:"passphrase"`
}

func (a *Account) GetKey() (*ecdsa.PrivateKey, error) {
	key, err := crypto.HexToECDSA(a.Key)
	if err != nil {
		return nil, err
	}

	return key, nil
}
