package handler

import (
	"errors"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/fx"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"net/http"

	"fmt"
)

var (
	supplierMap map[string]chan fx.Transaction
	resultChan  chan fx.ProcessResult
)

func init() {
	supplierMap = make(map[string]chan fx.Transaction)
	resultChan = make(chan fx.ProcessResult, 1)
}

func FxMintHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	companyId := m["companyId"].(string)
	if companyId == "" {
		render.Render(w, r, g.ErrBadRequest(errors.New("no companyId")))
		return
	}

	store := keychain.DefaultStore()
	_, err := store.GetAccount(companyId)
	if err != nil {
		log.Errorf("Query company failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if txType, err := fx.ParseType(m["txType"].(string)); err != nil {
		render.Render(w, r, g.ErrRender(err))
		return
	} else {
		tx := &fx.Transaction{TxType: txType}
		DistributeTask(tx, companyId)
		resp := g.NewSuccResponse("mint FX success")
		render.JSON(w, r, resp)
		return
	}

}

func DistributeTask(transaction *fx.Transaction, supplier string) {

	if supplierChain, ok := supplierMap[supplier]; ok {
		select {
		case supplierChain <- *transaction:
		default:
			fmt.Println("task for %v is full !", supplier)
		}

	} else {
		supplierChain := make(chan fx.Transaction, 5)
		supplierMap[supplier] = supplierChain
		supplierChain <- *transaction
		go fx.HandleTransaction(supplierChain, resultChan)
	}
}
