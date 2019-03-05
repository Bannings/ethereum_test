package handler

import (
	"errors"
	"net/http"

	"gitlab.com/ethereum_test/g"
	"gitlab.com/ethereum_test/keychain"

	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

type ctxKeyStore int

const (
	salt                 = "@cf#&"
	StoreKey ctxKeyStore = 0
)

func RegisterSupplierHandler(w http.ResponseWriter, r *http.Request) {
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
	log.Infof("companyId: %s", companyId)

	store := keychain.DefaultStore()
	account, err := store.GetAccount(companyId)
	if err != nil && err.Error() != fmt.Sprintf("account:%s not exist", companyId) {
		render.Render(w, r, g.ErrRender(err))
		return
	}
	if err == nil {
		render.JSON(w, r, g.NewSuccResponse(account, "Company already exist"))
		return
	}

	passphrase := companyId + salt
	acc, err := store.CreateAccount(passphrase)
	if err != nil {
		log.Errorf("Create account failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	err = store.StoreAccount(companyId, acc)
	if err != nil {
		log.Errorf("Store account failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	render.JSON(w, r, g.NewSuccResponse(acc))
}
