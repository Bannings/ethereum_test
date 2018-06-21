package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

const (
	pubPemStart = "-----BEGIN PUBLIC KEY-----"
	pubPemEnd   = "-----END PUBLIC KEY-----"
)

func AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	nodeName, ok1 := m["nodeName"]
	publicKey, ok2 := m["publicKey"]
	if !ok1 || !ok2 {
		log.Errorf("Bad params: %s", m)
		render.Render(w, r, g.ErrBadRequest(errors.New("no nodeName or publicKey")))
		return
	}
	log.Infof("node: %s", m)

	pubPem := publicKey.(string)
	if !strings.HasPrefix(pubPem, pubPemStart) {
		pubPem = fmt.Sprintf("%s\n%s\n%s", pubPemStart, pubPem, pubPemEnd)
	}

	delete(m, "nodeName")
	delete(m, "publicKey")
	b, _ := json.Marshal(m)
	s := string(b)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient().NodesTransactorSession(ctx).AddNode(nodeName.(string), pubPem, s)
	if err != nil {
		log.Errorf("Add Node to block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	var m g.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, g.ErrBadRequest(err))
		return
	}
	nodeName, ok := m["nodeName"]
	if !ok {
		log.Errorf("Bad params: %s", m)
		render.Render(w, r, g.ErrBadRequest(errors.New("no nodeName")))
		return
	}
	log.Infof("node: %s", m)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient().NodesTransactorSession(ctx).DeleteNode(nodeName.(string))
	if err != nil {
		log.Errorf("Delete Node from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := g.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func QueryNodeHandler(w http.ResponseWriter, r *http.Request) {
	node := r.URL.Query().Get("nodeName")
	if node == "" {
		render.Render(w, r, g.ErrBadRequest(errors.New("no nodeName")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pub, err := blockchain.DefaultClient().NodesCallerSession(ctx).GetNodeKey(node)
	log.Infof("publicKey: %v", pub)
	if err != nil {
		log.Errorf("Query node from block chain failed: %s", err.Error())
		render.Render(w, r, g.ErrRender(err))
		return
	}

	m := map[string]string{"nodeName": node, "publicKey": pub}
	log.Infof("ar: %+v", m)

	render.JSON(w, r, g.NewSuccResponse(m))
}
