package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/data"
	ex "gitlab.chainedfinance.com/chaincore/r2/error"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

func AddContractHandler(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("contract: %+v", m)

	contractNo := m["contractNo"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(contractNo)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if existed {
		render.Render(w, r, ex.ErrContractNoAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).AddContract(contractNo, s)
	if err != nil {
		log.Errorf("Add contract data to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)

}

func QueryContractHandler(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("contractNo")
	if no == "" {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No contractNo")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(no)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrContractNoNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, v, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryContractInfo(no)
	if err != nil {
		log.Errorf("Query contract from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var c data.M
	if err := json.Unmarshal([]byte(v), &c); err != nil {
		log.Errorf("Unmarshal contract data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	c["contractNo"] = no
	log.Infof("contract: %+v", c)

	resp := data.NewSuccResponse(map[string]data.M{"ContractInfo": c})
	render.JSON(w, r, resp)
}

func UpdateContractHander(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["contractNo"]
	if !ok {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No contractNo")))
		return
	}
	no := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(no)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrContractNoNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, c, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryContractInfo(no)
	if err != nil {
		log.Errorf("Query contract from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var contract data.M
	if err := json.Unmarshal([]byte(c), &contract); err != nil {
		log.Errorf("Unmarshal contract data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	log.Infof("contract from blockchain: %s", contract)

	for k, v := range m {
		contract[k] = v
	}
	log.Infof("merged contract: %+v", contract)

	b, _ := json.Marshal(contract)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).UpdateContractInfo(no, s)
	if err != nil {
		log.Errorf("update contract data to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}
