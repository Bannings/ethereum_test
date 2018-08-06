package handler

import (
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"gitlab.zhuronglink.com/chaincore/r2/keychain"
	"gitlab.zhuronglink.com/chaincore/r2/tokens"
	"net/http"

	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	supplierMap map[string]chan tokens.Transaction
	tradeTypes  []string
	state       []string
	db          *sql.DB
)

func init() {
	supplierMap = make(map[string]chan tokens.Transaction)
	tradeTypes = []string{"Payment", "Discount", "SplitToken", "MintToken", "Confirm"}
	state = []string{"Frozen", "Normal", "Historic"}
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	var trans tokens.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	log.Infof("Receive request:%v", body)
	err := json.Unmarshal(body, &trans)
	if err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	_, err = tokens.ParseType(trans.TxType)

	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	for _, token := range trans.Input {
		_, err = tokens.ParseState(token.State)
		if err != nil {
			resp := g.NewBadResponse("400", err.Error())
			render.JSON(w, r, resp)
			return
		}
	}
	for _, token := range trans.Output {
		_, err = tokens.ParseState(token.State)
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

func DistributeTask(transaction *tokens.Transaction) {
	var company string
	txType, _ := tokens.ParseType(transaction.TxType)

	if txType == tokens.MintToken {
		company = "cf"
	} else {
		company = transaction.Input[0].Owner
	}
	if supplierChain, ok := supplierMap[company]; ok {
		if len(supplierChain) > 0 {
			supplierChain <- *transaction
		} else {
			supplierChain <- *transaction
			go tokens.HandleTransaction(supplierChain)
		}
	} else {
		supplierChain := make(chan tokens.Transaction, 5)
		supplierMap[company] = supplierChain
		supplierChain <- *transaction
		go tokens.HandleTransaction(supplierChain)
	}
}

func saveTransaction(transaction *tokens.Transaction) error {
	input, err := json.Marshal(&transaction.Input)
	if err != nil {
		return err
	}

	output, err := json.Marshal(&transaction.Output)
	if err != nil {
		return err
	}

	stmt, err := tokens.DefaultDBConnection().Prepare("INSERT INTO transactions(deal_id, input, output,state) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = stmt.ExecContext(ctx, transaction.TxId, input, output, transaction.TxType)
	return err

}

func verifyTransaction(trans *tokens.Transaction) error {
	txType, err := tokens.ParseType(trans.TxType)
	if err != nil {
		return err
	}
	if txType == tokens.MintToken {
		if len(trans.Output) == 0 {
			return fmt.Errorf("Invalid transaction:%v, output not found")
		}
		for _, input := range trans.Input {
			_, err := keychain.DefaultStore().GetAccount(input.Owner)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		for _, inputToken := range trans.Input {
			_, err := keychain.DefaultStore().GetAccount(inputToken.Owner)
			if err != nil {
				return err
			}
			token, err := querTokenDetail(&inputToken.Id)
			if err != nil {
				return fmt.Errorf("Token:%v haven't been stored to blockchain or does't exist", inputToken.Id.Uint64())
			}
			state, err := tokens.ParseState(token.State)
			if err != nil {
				return err
			}
			if state == tokens.Frozen {
				return fmt.Errorf("Token:%v is frozen", inputToken.Id.Uint64())
			}
			if token.Owner != inputToken.Owner {
				return fmt.Errorf("The owner of token:%v is not %v", inputToken.Id.Uint64(), inputToken.Owner)
			}
			if token.Amount != inputToken.Amount {
				return fmt.Errorf("The amount of token:%v is wrong", inputToken.Id.Uint64())
			}
		}

		for _, outputToken := range trans.Output {
			_, err := keychain.DefaultStore().GetAccount(outputToken.Owner)
			if err != nil {
				return err
			}
			count := 0
			for _, inputToken := range trans.Input {
				if inputToken.Id.Uint64() == outputToken.ParentId.Uint64() {
					count++
				}
			}
			if count == 0 {
				return fmt.Errorf("Invalid transation output:%v ,can't find parentId for output", outputToken)
			}
		}
	}

	return nil
}
