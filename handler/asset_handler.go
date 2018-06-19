package handler

import (
	"errors"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/fx"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"net/http"

	"context"
	"encoding/json"
	"fmt"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"math/big"
	"time"
)

var (
	supplierMap map[string]chan fx.Transaction
	types       []string
)

func init() {
	supplierMap = make(map[string]chan fx.Transaction)
	types = []string{"Payment", "Discount", "SplitFX", "MintFX", "Confirm"}
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	var trans Transaction
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}

	tx, err := json.Marshal(m)
	fmt.Println(string(tx))
	if err != nil {
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	err = json.Unmarshal(tx, &trans)
	if err != nil {
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	_, err = fx.ParseType(trans.TxType)
	if err != nil {
		render.Render(w, r, g.ErrBadRequest(errors.New("Invalid TxType")))
		return
	}
	for _, token := range trans.Input {
		_, err = fx.ParseState(token.State)
		if err != nil {
			render.Render(w, r, g.ErrBadRequest(errors.New("Invalid input token state")))
			return
		}
	}
	for _, token := range trans.Output {
		_, err = fx.ParseState(token.State)
		if err != nil {
			render.Render(w, r, g.ErrBadRequest(errors.New("Invalid output token state")))
			return
		}
	}
	store := keychain.DefaultStore()
	exist, err := store.IsTransactionExist(trans.TxId)
	if exist {
		render.Render(w, r, g.ErrBadRequest(errors.New("Transaction id already exist")))
		return
	}
	err = saveTransaction(&trans)
	if err != nil {
		render.Render(w, r, g.ErrRender(err))
		return
	}
	resp := g.NewSuccResponse("Accept request success")
	render.JSON(w, r, resp)
	return

}

func DistributeTask(transaction *fx.Transaction) {
	var company string
	if transaction.TxType == fx.MintFX {
		company = "cf"
	} else {
		company = transaction.Input[0].Owner
	}
	if supplierChain, ok := supplierMap[company]; ok {
		if len(supplierChain) > 0 {
			supplierChain <- *transaction
		} else {
			supplierChain <- *transaction
			go fx.HandleTransaction(supplierChain)
		}
	} else {
		supplierChain := make(chan fx.Transaction, 5)
		supplierMap[company] = supplierChain
		supplierChain <- *transaction
		go fx.HandleTransaction(supplierChain)
	}
}

func saveTransaction(transaction *Transaction) error {
	conf := g.GetConfig()
	db, err := g.OpenDB(conf.DbConfig)
	if err != nil {
		return (err)
	}
	input, err := json.Marshal(&transaction.Input)
	if err != nil {
		return err
	}

	output, err := json.Marshal(&transaction.Output)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT INTO transactions(deal_id, input, output,state) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, transaction.TxId, input, output, transaction.TxType)
	return err

}

type Transaction struct {
	Id     uint    `json:"id"`
	Input  []Token `json:"input"`
	Output []Token `json:"output"`
	TxId   uint64  `json:"tx_id"`
	TxType string  `json:"tx_type"`
}

type Token struct {
	Id         big.Int `json:"id"`
	Amount     uint64  `json:"amount"`
	Owner      string  `json:"owner"` //company ID
	State      string  `json:"state"`
	ExpireTime int64   `json:"expire_time"`
}
