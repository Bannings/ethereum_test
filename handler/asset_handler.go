package handler

import (
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/fx"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"net/http"

	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	var trans fx.Transaction
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
	err = verifyTransaction(&trans)
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
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
	txType, _ := fx.ParseType(transaction.TxType)

	if txType == fx.MintFX {
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

func saveTransaction(transaction *fx.Transaction) error {
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

func verifyTransaction(trans *fx.Transaction) error {
	txType, err := fx.ParseType(trans.TxType)
	if err != nil {
		return err
	}
	if txType == fx.MintFX {
		for _, input := range trans.Input {
			_, err := keychain.DefaultStore().GetAccount(input.Owner)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		for _, input := range trans.Input {
			_, err := keychain.DefaultStore().GetAccount(input.Owner)
			if err != nil {
				return err
			}
			token, err := querFXDetail(&input.Id)
			if err != nil {
				errTxt := fmt.Sprintf("Token:%v haven't been stored to blockchain or does't exist", input.Id.Uint64())
				return errors.New(errTxt)
			}
			state, err := fx.ParseState(token.State)
			if err != nil {
				return err
			}
			if state == fx.Frozen {
				errTxt := fmt.Sprintf("Token:%v is frozen", input.Id.Uint64())
				return errors.New(errTxt)
			}
			if token.Owner != input.Owner {
				errTxt := fmt.Sprintf("The owner of token:%v is not %v", input.Id.Uint64(), input.Owner)
				return errors.New(errTxt)
			}
			if token.Amount != input.Amount {
				errTxt := fmt.Sprintf("The amount of token:%v is wrong", input.Id.Uint64())
				return errors.New(errTxt)
			}
		}
		for _, token := range trans.Output {
			_, err := keychain.DefaultStore().GetAccount(token.Owner)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
