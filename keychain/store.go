package keychain

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/eddyzhou/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Store struct {
	adminAccount Account

	client *Client
	db     *sql.DB

	mux sync.Mutex
}

func NewStore(adminAccount Account, rawUrl string, dbConfig g.DbConfig) (*Store, error) {
	c, err := newClient(adminAccount, rawUrl, 5*time.Second)
	if err != nil {
		log.Errorf("new client failed: %v", err)
		return nil, err
	}

	db, err := g.OpenDB(dbConfig)
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
	nonce := s.client.nonce
	tx := types.NewTransaction(nonce, to, amount, 100000, new(big.Int), nil)
	signTx, err := types.SignTx(tx, types.HomesteadSigner{}, key)
	//signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), s.adminPrivKey)
	if err != nil {
		return err
	}

	if err := s.client.SendTransaction(signTx); err != nil {
		return err
	}

	atomic.AddUint64(&s.client.nonce, 1)
	go func() {
		if _, err := bind.WaitMined(context.Background(), s.client.ethClient, tx); err != nil {
			atomic.StoreUint64(&s.client.nonce, nonce)
		}
	}()

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

	return Account{}, errors.New("account not exist")
}

func (s *Store) GetAdminAccount() Account {
	return s.adminAccount
}

func (s *Store) Close() error {
	err := s.db.Close()
	s.client.Close()
	return err
}
