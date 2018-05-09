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

func CreateArPayHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("arPay: %+v", m)

	payId := m["payId"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).HasPayment(payId)
	if err != nil {
		log.Errorf("Check payId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if existed {
		render.Render(w, r, g.ErrPayIdAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)
	log.Debug(s)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).PayByAR(payId, s)
	if err != nil {
		log.Errorf("Add arPay data to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)

}

func QueryArPayHandler(w http.ResponseWriter, r *http.Request) {
	payId := r.URL.Query().Get("payId")
	if payId == "" {
		render.Render(w, r, g.ErrBadRequest(errors.New("no payId")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).HasPayment(payId)
	if err != nil {
		log.Errorf("Check payId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrPayIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	v, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryPaymentTx(payId)
	if err != nil {
		log.Errorf("Query arPay from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var pay g.M
	if err := json.Unmarshal([]byte(v), &pay); err != nil {
		log.Errorf("Unmarshal arPay data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	pay["payId"] = payId
	log.Infof("arPay: %+v", pay)

	resp := g.NewSuccResponse(map[string]g.M{"txInfo": pay})
	render.JSON(w, r, resp)
}

func UpdateArPayHander(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["payId"]
	if !ok {
		render.Render(w, r, g.ErrBadRequest(errors.New("no payId")))
		return
	}
	payId := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).HasPayment(payId)
	if err != nil {
		log.Errorf("Check payId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrPayIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	p, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryPaymentTx(payId)
	if err != nil {
		log.Errorf("Query arPay from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var pay g.M
	if err := json.Unmarshal([]byte(p), &pay); err != nil {
		log.Errorf("Unmarshal arPay data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	log.Infof("arPay from blockchain: %s", pay)

	for k, v := range m {
		pay[k] = v
	}
	log.Infof("merged arPay: %+v", pay)

	b, _ := json.Marshal(pay)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).UpdatePaymentInfoOfAR(payId, s)
	if err != nil {
		log.Errorf("update arPay data to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}