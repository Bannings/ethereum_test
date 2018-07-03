package handler

import (
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/fx"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"net/http"

	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"io/ioutil"
	"math/big"
	"strings"
	"time"
)

var (
	supplierMap map[string]chan fx.Transaction
	tradeTypes  []string
	fxState     []string
	db          *sql.DB
)

func init() {
	supplierMap = make(map[string]chan fx.Transaction)
	tradeTypes = []string{"Payment", "Discount", "SplitFX", "MintFX", "Confirm"}
	fxState = []string{"Frozen", "Normal", "Historic"}
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	var trans Transaction
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, &trans)
	if err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	_, err = fx.ParseType(trans.TxType)
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	for _, token := range trans.Input {
		_, err = fx.ParseState(token.State)
		if err != nil {
			resp := g.NewBadResponse("400", err.Error())
			render.JSON(w, r, resp)
			return
		}
	}
	for _, token := range trans.Output {
		_, err = fx.ParseState(token.State)
		if err != nil {
			resp := g.NewBadResponse("400", err.Error())
			render.JSON(w, r, resp)
			return
		}
	}
	store := keychain.DefaultStore()
	exist, err := store.IsTransactionExist(trans.TxId)
	if exist {
		resp := g.NewBadResponse("400", "Transaction id already exist")
		render.JSON(w, r, resp)
		return
	}
	result, companyId := verifyTransaction(&trans)
	if !result {
		errText := "company id" + *companyId + " not exist"
		resp := g.NewBadResponse("400", errText)
		render.JSON(w, r, resp)
		return
	}
	err = saveTransaction(&trans)
	if err != nil {
		log.Errorf("%v", err)
		resp := g.NewBadResponse("500", err.Error())
		render.JSON(w, r, resp)
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
	input, err := json.Marshal(&transaction.Input)
	if err != nil {
		return err
	}

	output, err := json.Marshal(&transaction.Output)
	if err != nil {
		return err
	}

	stmt, err := fx.DefaultDBConnection().Prepare("INSERT INTO transactions(deal_id, input, output,state) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, transaction.TxId, input, output, transaction.TxType)
	return err

}

func verifyTransaction(trans *Transaction) (bool, *string) {
	for _, token := range trans.Input {
		_, err := keychain.DefaultStore().GetAccount(token.Owner)
		if err != nil {
			return false, &token.Owner
		}
	}
	if strings.ToLower(trans.TxType) == "mintfx" {
		return true, nil
	}
	for _, token := range trans.Output {
		_, err := keychain.DefaultStore().GetAccount(token.Owner)
		if err != nil {
			return false, &token.Owner
		}
	}
	return true, nil
}

type Transaction struct {
	Id     uint    `json:"id"`
	Input  []Token `json:"input"`
	Output []Token `json:"output"`
	TxId   string  `json:"tx_id"`
	TxType string  `json:"tx_type"`
}

type Token struct {
	Id         big.Int `json:"id"`
	ParentId   big.Int `json:"parentId"`
	Amount     uint64  `json:"amount"`
	Owner      string  `json:"owner"` //company ID
	State      string  `json:"state"`
	ExpireTime int64   `json:"expire_time"`
}
