package keychain

import (
	"crypto/ecdsa"
	"errors"

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

func GetAccount(privKey, passphrase string) (Account, error) {
	if privKey == "" {
		return Account{}, errors.New("err: empty private key")
	}

	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return Account{}, err
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	return Account{Address: address, Key: privKey, Passphrase: passphrase}, nil
}
