package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

func AddContractHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("contract: %+v", m)

	contractNo := m["contractNo"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(contractNo)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if existed {
		render.Render(w, r, g.ErrContractNoAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).AddContract(contractNo, s)
	if err != nil {
		log.Errorf("Add contract data to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)

}

func QueryContractHandler(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("contractNo")
	if no == "" {
		render.Render(w, r, g.ErrBadRequest(errors.New("no contractNo")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(no)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrContractNoNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, v, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryContractInfo(no)
	if err != nil {
		log.Errorf("Query contract from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var c g.M
	if err := json.Unmarshal([]byte(v), &c); err != nil {
		log.Errorf("Unmarshal contract data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	c["contractNo"] = no
	log.Infof("contract: %+v", c)

	resp := g.NewSuccResponse(map[string]g.M{"ContractInfo": c})
	render.JSON(w, r, resp)
}

func UpdateContractHander(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["contractNo"]
	if !ok {
		render.Render(w, r, g.ErrBadRequest(errors.New("no contractNo")))
		return
	}
	no := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckContractNo(no)
	if err != nil {
		log.Errorf("Check contractNo failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrContractNoNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, c, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryContractInfo(no)
	if err != nil {
		log.Errorf("Query contract from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var contract g.M
	if err := json.Unmarshal([]byte(c), &contract); err != nil {
		log.Errorf("Unmarshal contract data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
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
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}
