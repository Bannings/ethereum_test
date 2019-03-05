package keychain

import (
	"context"
	"database/sql"
	"encoding/hex"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.com/ethereum_test/g"

	"fmt"
	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	defaultStore *Store
	once         sync.Once
)

func DefaultStore() *Store {
	once.Do(func() {
		conf := g.GetConfig()
		bConf := conf.BlockchainConfig
		acc, err := GetAccount(bConf.AdminKey, bConf.AdminPassphrase)
		if err != nil {
			panic(err)
		}
		store, err := NewStore(acc, bConf.RawUrl, conf.DbConfig)
		if err != nil {
			panic(err)
		}
		defaultStore = store
	})
	return defaultStore
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return big.NewInt(1000000000000000000)
}

type Store struct {
	adminAccount Account
	adminClient  *AccountClient
	db           *sql.DB

	mux sync.Mutex
}

func NewStore(adminAccount Account, rawUrl string, dbConfig g.DbConfig) (*Store, error) {
	c, err := NewAccountClient(adminAccount, rawUrl, 5*time.Second)
	if err != nil {
		log.Errorf("new client failed: %v", err)
		return nil, err
	}

	db, err := g.OpenDB(dbConfig)
	if err != nil {
		return nil, err
	}

	return &Store{adminAccount: adminAccount, adminClient: c, db: db}, nil
}

func (s *Store) CreateAccount(passphrase string) (Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return Account{}, err
	}

	keyHex := hex.EncodeToString(crypto.FromECDSA(key))
	log.Debugf("key: %s", keyHex)
	//addr, err := s.adminClient.PersonalImportRawKey(keyHex, passphrase)
	//if err != nil {
	//	log.Errorf("personal_importRawKey %s failed: %v", keyHex, err)
	//	return Account{}, err
	//}
	//log.Debugf("--- addr: %s", addr)

	//if _, err := s.client.PersonalUnlockAccount(s.adminAccount.Address.Hex(), s.adminAccount.Passphrase, 30); err != nil {
	//	log.Errorf("Unlock admin account failed: %v", err)
	//	return Account{}, err
	//}

	address := crypto.PubkeyToAddress(key.PublicKey)

	v := new(big.Int)
	v = v.Mul(Eth1(), big.NewInt(50))
	if err := s.TransferEther(address, v); err != nil {
		log.Errorf("transfer Ether failed: %v", err)
		return Account{}, err
	}

	return Account{address, keyHex, passphrase}, nil
}

func (s *Store) TransferEther(to common.Address, amount *big.Int) error {
	key, err := crypto.HexToECDSA(s.adminAccount.Key)
	if err != nil {
		return err
	}

	s.mux.Lock()
	defer s.mux.Unlock()
	nonce := s.adminClient.Nonce()
	log.Debugf("nonce: %v", nonce)
	tx := types.NewTransaction(nonce, to, amount, 100000, new(big.Int), nil)
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, key)
	//signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), s.adminPrivKey)
	if err != nil {
		return err
	}

	if err := s.adminClient.SendTransaction(signTx); err != nil {
		return err
	}

	s.adminClient.IncrNonce()
	go func() {
		if _, err := bind.WaitMined(context.Background(), s.adminClient.EthClient, tx); err != nil {
			log.Errorf("tx execute failed: %v", err)
			atomic.StoreUint64(s.adminClient.nonce, nonce)
		}
	}()
	log.Debugf("now nonce: %v", s.adminClient.Nonce())

	return nil
}

func (s *Store) StoreAccount(companyID string, account Account) error {
	stmt, err := s.db.Prepare("INSERT INTO accounts(firm_id, address, priv_key, passphrase) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, companyID, account.Address.Hex(), account.Key, account.Passphrase)
	return err
}

func (s *Store) GetAccount(companyID string) (Account, error) {
	rows, err := s.db.Query("select address, priv_key, passphrase from accounts where firm_id = ? order by id desc limit 1", companyID)
	if err != nil {
		log.Warnf("fail to query firm keystore account: %v", err)
		return Account{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var address, key, passphrase string
		rows.Scan(&address, &key, &passphrase)
		return Account{Address: common.HexToAddress(address), Key: key, Passphrase: passphrase}, nil
	}
	return Account{}, fmt.Errorf("account:%s not exist", companyID)
}

func (s *Store) IsAccountExist(companyID string) (bool, error) {
	var result bool
	rows, err := s.db.Query("select isnull ((select 1 from accounts  where firm_id = ? limit 1))", companyID)
	if err != nil {
		log.Warnf("fail to query firm company: %v", err)
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count == 0 {
			result = true
		} else {
			result = false
		}
	}
	return result, nil
}

func (s *Store) IsTransactionExist(TxId string) (bool, error) {
	var result bool
	rows, err := s.db.Query("select isnull ((select 1 from transactions where deal_id = ? limit 1))", TxId)
	if err != nil {
		log.Warnf("fail to query transaction id : %v", err)
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		var count int
		rows.Scan(&count)
		if count == 0 {
			result = true
		} else {
			result = false
		}
	}
	return result, nil
}

func (s *Store) GetAdminAccount() Account {
	return s.adminAccount
}

func (s *Store) GetAdminClient() *AccountClient {
	return s.adminClient
}

func (s *Store) Close() error {
	err := s.db.Close()
	s.adminClient.Close()
	return err
}
