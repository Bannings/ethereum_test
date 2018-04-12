package keychain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/boltdb/bolt"
	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	accoutBucket = "accounts"
)

type Store struct {
	adminAccount Account

	client *Client
	db     *bolt.DB // TODO: boltdb to mysql

	mux sync.Mutex
}

func NewStore(adminAccount Account, rawUrl, dir string) (*Store, error) {
	c, err := newClient(adminAccount, rawUrl, 2*time.Second)
	if err != nil {
		return nil, err
	}

	db, err := bolt.Open(filepath.Join(dir, "accounts.db"), 0600, nil)
	if err != nil {
		return nil, err
	}

	return &Store{adminAccount: adminAccount, client: c, db: db}, nil
}

func (s *Store) CreateAccount(passphrase string) (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, err
	}

	keyHex := hex.EncodeToString(crypto.FromECDSA(key))
	addr, err := s.client.PersonalImportRawKey(keyHex, passphrase)
	if err != nil {
		return Account{}, err
	}
	log.Debugf("--- addr: %s", addr)

	//if _, err := s.client.PersonalUnlockAccount(s.adminAccount.Address.Hex(), s.adminAccount.Passphrase, 30); err != nil {
	//	log.Errorf("Unlock admin account failed: %v", err)
	//	return Account{}, err
	//}

	address := crypto.PubkeyToAddress(key.PublicKey)

	v := new(big.Int)
	v = v.Mul(Eth1(), big.NewInt(1000))
	if err := s.transferEther(address, v); err != nil {
		log.Errorf("transfer Ether failed: %v", err)
		return Account{}, err
	}

	return Account{address, keyHex, passphrase}, nil
}

func (s *Store) transferEther(to common.Address, amount *big.Int) error {
	key, err := crypto.HexToECDSA(s.adminAccount.Key)
	if err != nil {
		return err
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	tx := types.NewTransaction(s.client.nonce, to, amount, 100000, new(big.Int), nil)
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, key)
	//signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), s.adminPrivKey)
	if err != nil {
		return err
	}

	if err := s.client.SendTransaction(signTx); err != nil {
		return err
	}

	atomic.AddUint64(&s.client.nonce, 1)
	return nil
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

func (s *Store) GetAdminAccount() Account {
	return s.adminAccount
}

func (s *Store) Close() error {
	err := s.db.Close()
	s.client.Close()
	return err
}
