package handler

import (
	"context"
	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"gitlab.zhuronglink.com/chaincore/r2/tokens"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func TokenManagementHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	tokenId := m["tokenId"].(string)
	state := m["state"].(string)
	if tokenId == "" && state == "" {
		resp := g.NewBadResponse("400", "Missing necessary parameter:tokenId")
		render.JSON(w, r, resp)
		return
	}
	_, err := tokens.ParseState(state)
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}

	id, err := strconv.ParseInt(tokenId, 10, 64)
	if err != nil {
		resp := g.NewBadResponse("400", "Invalid tokenId:tokenId must be int type")
		render.JSON(w, r, resp)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = tokenVerify(big.NewInt(id), ctx)
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	token := tokens.Token{Id: *big.NewInt(id), State: state}
	tx := tokens.Transaction{Input: []tokens.Token{token}, TxType: "Setting"}
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
