package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/keychain"
	"gitlab.chainedfinance.com/chaincore/r2/tokens"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func QueryFXHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}

	fxId := m["tokenId"].(string)
	if fxId == "" {
		resp := g.NewBadResponse("400", "no tokenId found")
		render.JSON(w, r, resp)
		return
	}
	id, err := strconv.ParseInt(fxId, 10, 64)
	if err != nil {
		resp := g.NewBadResponse("400", "invalid tokenId")
		render.JSON(w, r, resp)
		return
	}
	token, err := querFXDetail(big.NewInt(id))
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	render.JSON(w, r, g.NewSuccResponse(token))
}

func querFXDetail(fxID *big.Int) (*tokens.Token, error) {
	conf := g.GetConfig()
	bConf := conf.BlockchainConfig
	store := keychain.DefaultStore()
	adminClient, err := blockchain.NewTokenClient(store.GetAdminClient(), store.GetAdminAccount(), bConf.ContractAddrs)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := adminClient.CallWithFXStorageCaller(ctx).Existed(fxID)
	if !existed {
		errTxt := fmt.Sprintf("FX:%s is not exist", fxID.Uint64())
		return nil, errors.New(errTxt)
	}

	owner, err := adminClient.CallWithERC721TokenCaller(ctx).OwnerOf(fxID)
	var company string
	err = tokens.DefaultDBConnection().QueryRow("SELECT firm_id FROM accounts WHERE address = ?", owner.String()).Scan(&company)
	if err != nil {
		return nil, err
	}
	properties, err := adminClient.CallWithFXStorageCaller(ctx).GetProperties(fxID, false)
	if err != nil {
		return nil, err
	}
	stateID, err := adminClient.CallWithFXLockerCaller(ctx).GetState(fxID)
	if err != nil {
		return nil, err
	}
	state := fxState[stateID.Uint64()]
	expire, err := adminClient.CallWithFXLockerCaller(ctx).GetExpire(fxID)
	if err != nil {
		return nil, err
	}
	token := &tokens.Token{Id: *fxID, Amount: properties.Value.Uint64(), ParentId: *properties.CreateBy, State: state, Owner: company, ExpireTime: expire.Int64()}
	return token, nil
}
