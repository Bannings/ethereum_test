package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eddyzhou/log"
	"github.com/go-chi/render"
	"gitlab.chainedfinance.com/chaincore/r2/g"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

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

		var m g.M
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
		return "", errors.New("not supported method")
	}
}

func OnlyCF(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		node := r.Header.Get("nodeName")
		if node != "cf" && node != "CF" {
			render.Render(w, r, g.ErrBadRequest(errors.New("node err: Only CF")))
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
