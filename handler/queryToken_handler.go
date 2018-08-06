package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.zhuronglink.com/chaincore/r2/blockchain"
	"gitlab.zhuronglink.com/chaincore/r2/g"
	"gitlab.zhuronglink.com/chaincore/r2/keychain"
	"gitlab.zhuronglink.com/chaincore/r2/tokens"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

func QueryTokenHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed: %s", err.Error())
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}

	tokenId := m["tokenId"].(string)
	if tokenId == "" {
		resp := g.NewBadResponse("400", "no tokenId found")
		render.JSON(w, r, resp)
		return
	}
	id, err := strconv.ParseInt(tokenId, 10, 64)
	if err != nil {
		resp := g.NewBadResponse("400", "invalid tokenId")
		render.JSON(w, r, resp)
		return
	}
	token, err := querTokenDetail(big.NewInt(id))
	if err != nil {
		resp := g.NewBadResponse("400", err.Error())
		render.JSON(w, r, resp)
		return
	}
	render.JSON(w, r, g.NewSuccResponse(token))
}

func querTokenDetail(tokenID *big.Int) (*tokens.Token, error) {
	conf := g.GetConfig()
	bConf := conf.BlockchainConfig
	store := keychain.DefaultStore()
	adminClient, err := blockchain.NewTokenClient(store.GetAdminClient(), store.GetAdminAccount(), bConf.ContractAddrs)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	existed, err := adminClient.CallWithStorageCaller(ctx).Existed(tokenID)
	if !existed {
		errTxt := fmt.Sprintf("Token:%v is not exist", tokenID.Uint64())
		return nil, errors.New(errTxt)
	}

	owner, err := adminClient.CallWithERC721TokenCaller(ctx).OwnerOf(tokenID)
	var company string
	err = tokens.DefaultDBConnection().QueryRow("SELECT firm_id FROM accounts WHERE address = ?", owner.String()).Scan(&company)
	if err != nil {
		return nil, err
	}
	properties, err := adminClient.CallWithStorageCaller(ctx).GetProperties(tokenID, false)
	if err != nil {
		return nil, err
	}
	stateID, err := adminClient.CallWithLockerCaller(ctx).GetState(tokenID)
	if err != nil {
		return nil, err
	}
	state := state[stateID.Uint64()]
	expire, err := adminClient.CallWithLockerCaller(ctx).GetExpire(tokenID)
	if err != nil {
		return nil, err
	}
	token := &tokens.Token{Id: *tokenID, Amount: properties.Value.Uint64(), ParentId: *properties.CreateBy, State: state, Owner: company, ExpireTime: expire.Int64()}
	return token, nil
}
