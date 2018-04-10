package middleware

import (
	"context"
	"net/http"
)

type ctxKeyVersion string

const VersionKey ctxKeyVersion = "api.version"

func ApiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), VersionKey, version))
			next.ServeHTTP(w, r)
		})
	}
}
