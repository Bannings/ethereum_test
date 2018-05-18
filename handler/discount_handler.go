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

func CreateDiscountHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("discount: %+v", m)

	discountId := m["discountId"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient().StoreCallerSession(ctx).HasDiscount(discountId)
	if err != nil {
		log.Errorf("Check discountId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if existed {
		render.Render(w, r, g.ErrDiscountIdAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient().StoreTransactorSession(ctx).DiscountByAR(discountId, s)
	if err != nil {
		log.Errorf("Add discount data to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)

}

func QueryDiscountHandler(w http.ResponseWriter, r *http.Request) {
	discountId := r.URL.Query().Get("discountId")
	if discountId == "" {
		render.Render(w, r, g.ErrBadRequest(errors.New("no discountId")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient().StoreCallerSession(ctx).HasDiscount(discountId)
	if err != nil {
		log.Errorf("Check discountId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrDiscountIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	v, err := blockchain.DefaultClient().StoreCallerSession(ctx).QueryDiscountTx(discountId)
	if err != nil {
		log.Errorf("Query discount from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var discount g.M
	if err := json.Unmarshal([]byte(v), &discount); err != nil {
		log.Errorf("Unmarshal discount data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	discount["discountId"] = discountId
	log.Infof("discount: %+v", discount)

	resp := g.NewSuccResponse(map[string]g.M{"txInfo": discount})
	render.JSON(w, r, resp)
}

func UpdateDiscountHander(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["discountId"]
	if !ok {
		render.Render(w, r, g.ErrBadRequest(errors.New("no discountId")))
		return
	}
	discountId := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient().StoreCallerSession(ctx).HasDiscount(discountId)
	if err != nil {
		log.Errorf("Check discountId failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, g.ErrDiscountIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	d, err := blockchain.DefaultClient().StoreCallerSession(ctx).QueryDiscountTx(discountId)
	if err != nil {
		log.Errorf("Query discount from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	var discount g.M
	if err := json.Unmarshal([]byte(d), &discount); err != nil {
		log.Errorf("Unmarshal arPay data failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	log.Infof("discount from blockchain: %v", discount)

	for k, v := range m {
		discount[k] = v
	}
	log.Infof("merged discount: %+v", discount)

	b, _ := json.Marshal(discount)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient().StoreTransactorSession(ctx).UpdateDiscountInfoOfAR(discountId, s)
	if err != nil {
		log.Errorf("update discount data to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}
