package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"fmt"
	"strings"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/data"
	ex "gitlab.chainedfinance.com/chaincore/r2/error"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

const (
	pubPemStart = "-----BEGIN PUBLIC KEY-----"
	pubPemEnd   = "-----END PUBLIC KEY-----"
)

func AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	nodeName, ok1 := m["nodeName"]
	publicKey, ok2 := m["publicKey"]
	if !ok1 || !ok2 {
		log.Errorf("Bad params: %s", m)
		render.Render(w, r, ex.ErrBadRequest(errors.New("No nodeName or publicKey")))
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
	tx, err := blockchain.DefaultClient.NodesTransactorSession(ctx).AddNode(nodeName.(string), pubPem, s)
	if err != nil {
		log.Errorf("Add Node to block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	var m data.M
	if err := render.Bind(r, &m); err != nil {
		log.Errorf("Unmarshal request failed. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}
	nodeName, ok := m["nodeName"]
	if !ok {
		log.Errorf("Bad params: %s", m)
		render.Render(w, r, ex.ErrBadRequest(errors.New("No nodeName")))
		return
	}
	log.Infof("node: %s", m)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tx, err := blockchain.DefaultClient.NodesTransactorSession(ctx).DeleteNode(nodeName.(string))
	if err != nil {
		log.Errorf("Delete Node from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	txHash := tx.Hash().Hex()

	resp := data.NewSuccResponse(map[string]string{"txHash": txHash})
	render.JSON(w, r, resp)
}

func QueryNode(w http.ResponseWriter, r *http.Request) {
	node := r.URL.Query().Get("nodeName")
	if node == "" {
		render.Render(w, r, ex.ErrBadRequest(errors.New("No nodeName")))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, v, err := blockchain.DefaultClient.NodesCallerSession(ctx).GetNodeKey(node)
	if err != nil {
		log.Errorf("Query node from block chain failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}

	var m data.M
	if err := json.Unmarshal([]byte(v), &m); err != nil {
		log.Errorf("Unmarshal node data failed: %s", err.Error())
		render.Render(w, r, ex.ErrRender(err))
		return
	}
	log.Infof("ar: %+v", m)

	resp := data.NewSuccResponse(m)
	render.JSON(w, r, resp)
}
