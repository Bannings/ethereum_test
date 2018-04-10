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
	"gitlab.chainedfinance.com/chaincore/keychain"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

var AccountStore keychain.Store

func CreateArHandler(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("ar: %+v", m)

	arId := m["arId"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckARId(arId)
	log.Debugf("existed: %v", existed)
	if err != nil {
		log.Errorf("Check arId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if existed {
		log.Errorf("arId already exist: %s", arId)
		render.Render(w, r, ex.ErrArIdAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)
	log.Infof("ar info: %s", s)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).CreateAR(arId, s)
	if err != nil {
		log.Errorf("Add AR data to block chain failed: %v", err)
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func QueryArHandler(w http.ResponseWriter, r *http.Request) {
	arId := r.URL.Query().Get("arId")
	if arId == "" {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No arId")))
		return
	}
	log.Infof("arId: %v", arId)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckARId(arId)
	if err != nil {
		log.Errorf("Check arId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrArIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, v, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryArInfo(arId)
	if err != nil {
		log.Errorf("Query ar from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var m data.M
	if err := json.Unmarshal([]byte(v), &m); err != nil {
		log.Errorf("Unmarshal ar data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	m["arId"] = arId
	log.Infof("ar: %+v", m)

	resp := data.NewSuccResponse(map[string]data.M{"ArInfo": m})
	render.JSON(w, r, resp)
}

func UpdateArHander(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["arId"]
	if !ok {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No arId")))
		return
	}
	arId := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckARId(arId)
	if err != nil {
		log.Errorf("Check arId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrArIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, info, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryArInfo(arId)
	if err != nil {
		log.Errorf("Query ar from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var ar data.M
	if err := json.Unmarshal([]byte(info), &ar); err != nil {
		log.Errorf("Unmarshal ar data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	log.Infof("ar from blockchain: %+v", ar)

	for k, v := range m {
		ar[k] = v
	}
	log.Infof("merged ar: %+v", ar)

	b, _ := json.Marshal(ar)
	s := string(b)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).UpdateARInfo(arId, s)
	if err != nil {
		log.Errorf("update ar data to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}
