package handler

import (
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/fx"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"net/http"

	"context"
	"encoding/json"
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
	var trans fx.Transaction
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}

	tx, err := json.Marshal(m)
	err = json.Unmarshal(tx, trans)
	if err != nil {
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	err = saveTransaction(&trans)
	if err != nil {
		render.Render(w, r, g.ErrRender(err))
		return
	}
	//DistributeTask(&trans)
	resp := g.NewAcceptResponse("Accept request success")
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

func saveTransaction(transaction *fx.Transaction) error {
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
	_, err = stmt.ExecContext(ctx, transaction.TxId, input, output, types[transaction.TxType])
	return err

}
