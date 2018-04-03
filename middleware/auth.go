package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"
	"gitlab.chainedfinance.com/chaincore/r2/data"
	ex "gitlab.chainedfinance.com/chaincore/r2/error"
	"gitlab.chainedfinance.com/chaincore/r2/sig"

	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
)

func Auth() func(http.Handler) http.Handler {
	a := &author{
		client: blockchain.DefaultClient,
	}
	fn := func(h http.Handler) http.Handler {
		a.h = h
		return a
	}
	return fn
}

type author struct {
	h      http.Handler
	client blockchain.Client
}

func (a *author) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("sig")
	node := r.Header.Get("nodeName")
	if signature == "" {
		log.Error("Auth: no sig Header")
		render.Render(w, r, ex.ErrBadRequest(errors.New("No sig Header")))
		return
	}
	if node == "" {
		log.Error("Auth: no nodeName Header")
		render.Render(w, r, ex.ErrBadRequest(errors.New("No nodeName Header")))
		return
	}

	baseStr, err := baseString(r)
	log.Debugf("baseString: %s", baseStr)
	if err != nil {
		log.Errorf("Auth: bad request. %s", err.Error())
		render.Render(w, r, ex.ErrBadRequest(err))
		return
	}

	pubPEM, err := sig.GetKey(a.client, node)
	if err != nil {
		log.Errorf("Auth: node err: %s", node)
		render.Render(w, r, ex.ErrInvalidNode)
		return
	}

	var errSig error
	if errSig = sig.VerifySHA256RSA(pubPEM, []byte(baseStr), signature); errSig != nil {
		if pubPEM, errSig = sig.FetchKey(a.client, node); errSig == nil {
			errSig = sig.VerifySHA256RSA(pubPEM, []byte(baseStr), signature)
		}
	}
	if errSig != nil {
		log.Errorf("Auth: verify sig failed. %s", errSig.Error())
		render.Render(w, r, ex.ErrVerifySignature)
		return
	}

	a.h.ServeHTTP(w, r)
}

func baseString(r *http.Request) (string, error) {
	switch r.Method {
	case http.MethodGet:
		query := r.URL.Query()
		sortedKv := make([]string, 0, len(query))
		for k, v := range query {
			sortedKv = append(sortedKv, fmt.Sprintf("%s%s", k, strings.Join(v, ",")))
		}
		sort.Strings(sortedKv)
		return strings.Join(sortedKv, ""), nil

	case http.MethodPost:
		if render.GetRequestContentType(r) != render.ContentTypeJSON {
			return "", errors.New("ContentType err")
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return "", err
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		var m data.M
		if err := json.Unmarshal(body, &m); err != nil {
			return "", err
		}

		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		sortedKv := make([]string, 0, len(m))
		for _, k := range keys {
			v := m[k]
			log.Debugf("%s=%s", k, v)
			sortedKv = append(sortedKv, fmt.Sprintf("%s%s", k, v))
		}
		return strings.Join(sortedKv, ""), nil

	default:
		return "", errors.New("Not supported method")
	}
}


func OnlyCF(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		node := r.Header.Get("nodeName")
		if node != "cf" && node != "CF"{
			render.Render(w, r, ex.ErrBadRequest(errors.New("node err: Only CF")))
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
