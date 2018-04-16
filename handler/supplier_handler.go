package handler

import (
	"errors"
	"net/http"

	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"

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

	ctx := r.Context()
	store := ctx.Value(StoreKey).(keychain.Store)

	passphrase := companyId + salt
	acc, err := store.CreateAccount(passphrase)
	if err != nil {
		log.Errorf("Create account failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	if err := store.StoreAccount(companyId, acc); err != nil {
		log.Errorf("Store account failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	render.JSON(w, r, g.NewSuccResponse(acc))
}
