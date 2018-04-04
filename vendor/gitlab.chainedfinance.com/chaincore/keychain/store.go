package keychain

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"path/filepath"
	"time"
)

const (
	accoutBucket = "accounts"
)

type Store struct {
	ethClient *rpc.Client
	db        *bolt.DB
}

func NewStore(rawUrl, dir string) (*Store, error) {
	c, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, err
	}

	db, err := bolt.Open(filepath.Join(dir, "my.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	return &Store{c, db}, nil
}

func (s *Store) CreateAccount(passphrase string) (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, err
	}

	keyHex := hex.EncodeToString(crypto.FromECDSA(key))
	var addr string
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := s.ethClient.CallContext(ctx, &addr, "personal_importRawKey", keyHex, passphrase); err != nil {
		return Account{}, err
	}
	log.Debugf("--- addr: %s", addr)

	address := crypto.PubkeyToAddress(key.PublicKey)
	return Account{address, keyHex, passphrase}, nil
}

func (s *Store) StoreAccount(companyID string, account Account) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(accoutBucket))
		if err != nil {
			return err
		}

		buf, err := json.Marshal(account)
		if err != nil {
			return err
		}

		return b.Put([]byte(companyID), buf)
	})
}

func (s *Store) GetAccount(companyID string) (Account, error) {
	var data []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(accoutBucket))
		data = b.Get([]byte(companyID))
		if data == nil {
			return errors.New("Account not exist")
		}
		return nil
	})

	if err != nil {
		return Account{}, err
	}

	var ac Account
	if err := json.Unmarshal(data, &ac); err != nil {
		return Account{}, err
	}

	return ac, nil
}

func (s *Store) Close() error {
	err := s.db.Close()
	s.ethClient.Close()
	return err
}
