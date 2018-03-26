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

func AddCompanyHandler(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("company: %+v", m)

	companyId := m["companyId"].(string)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckCompanyId(companyId)
	if err != nil {
		log.Errorf("Check companyId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if existed {
		render.Render(w, r, ex.ErrCompanyIdAlreadyExist)
		return
	}

	b, _ := json.Marshal(m)
	s := string(b)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).AddCompany(companyId, s)
	if err != nil {
		log.Errorf("Add company data to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func QueryCompanyHandler(w http.ResponseWriter, r *http.Request) {
	companyId := r.URL.Query().Get("companyId")
	if companyId == "" {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No companyId")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckCompanyId(companyId)
	if err != nil {
		log.Errorf("Check companyId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrCompanyIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, v, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryCompanyInfo(companyId)
	if err != nil {
		log.Errorf("Query company from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var company data.M
	if err := json.Unmarshal([]byte(v), &company); err != nil {
		log.Errorf("Unmarshal company data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	company["companyId"] = companyId
	log.Infof("company: %+v", company)

	resp := data.NewSuccResponse(map[string]data.M{"CompanyInfo": company})
	render.JSON(w, r, resp)
}

func UpdateCompanyHander(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	log.Infof("params: %v", m)

	v, ok := m["companyId"]
	if !ok {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No companyId")))
		return
	}
	companyId := v.(string)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := blockchain.DefaultClient.StoreCallerSession(ctx).CheckCompanyId(companyId)
	if err != nil {
		log.Errorf("Check companyId failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	if !existed {
		render.Render(w, r, ex.ErrCompanyIdNotExist)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, c, err := blockchain.DefaultClient.StoreCallerSession(ctx).QueryCompanyInfo(companyId)
	if err != nil {
		log.Errorf("Query company from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var company data.M
	if err := json.Unmarshal([]byte(c), &company); err != nil {
		log.Errorf("Unmarshal company data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	log.Infof("company from blockchain: %+v", company)

	for k, v := range m {
		company[k] = v
	}
	log.Infof("merged company: %+v", company)

	b, _ := json.Marshal(company)
	s := string(b)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.StoreTransactorSession(ctx).UpdateCompanyInfo(companyId, s)
	if err != nil {
		log.Errorf("update company data to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}
