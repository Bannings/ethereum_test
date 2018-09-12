package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"gitlab.zhuronglink.com/chaincore/r2/keychain"
	"gitlab.zhuronglink.com/chaincore/r2/tokens"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"
)

func TokenManagementHandler(w http.ResponseWriter, r *http.Request) {
	var management tokens.Management
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &management)
	if err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	store := keychain.DefaultStore()
	exist, err := store.IsTransactionExist(management.TxId)
	if exist {
		resp := g.NewBadResponse("400", "Transaction id already exist")
		render.JSON(w, r, resp)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, token := range management.Tokens {
		_, err = tokens.ParseState(token.State)
		if err != nil {
			resp := g.NewBadResponse("400", err.Error())
			render.JSON(w, r, resp)
			return
		}

		err = tokenVerify(&token.Id, ctx)
		if err != nil {
			resp := g.NewBadResponse("400", err.Error())
			render.JSON(w, r, resp)
			return
		}
	}

	tx := tokens.Transaction{TxId: management.TxId, Input: management.Tokens, TxType: management.Type}
	err = saveTransaction(&tx)
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}

	resp := g.NewSuccResponse("Accept request success")
	render.JSON(w, r, resp)
	return
}

func tokenVerify(tokenID *big.Int, ctx context.Context) error {
	existed, err := tokenExistOrNot(tokenID, ctx)
	if err != nil {
		return err
	}
	if !existed {
		return fmt.Errorf("Token id:%v not exist")
	}
	return nil
}
