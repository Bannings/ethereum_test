package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"gitlab.chainedfinance.com/chaincore/r2/g"

	"github.com/eddyzhou/log"
	"github.com/getsentry/raven-go"
	"github.com/go-chi/render"
)

const (
	MAXSTACKSIZE = 4096
)

type Recoverer struct {
	h            http.Handler
	monitor      *Monitor
	sentryClient *raven.Client
}

func NewRecoverer(m *Monitor, sentryDSN string) (*Recoverer, error) {
	client, err := raven.New(sentryDSN)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	r := &Recoverer{
		monitor:      m,
		sentryClient: client,
	}

	return r, nil
}

func Recovery(m *Monitor, sentryDSN string) func(http.Handler) http.Handler {
	r, err := NewRecoverer(m, sentryDSN)
	if err != nil {
		panic(err)
	}

	fn := func(h http.Handler) http.Handler {
		r.h = h
		return r
	}

	return fn
}

func (rec *Recoverer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			if rw.Header().Get("Content-Type") == "" {
				rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
			}

			rw.WriteHeader(http.StatusInternalServerError)

			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]

			f := "PANIC: %s\n%s"
			log.Errorf(f, err, stack)

			var desc string
			switch rval := err.(type) {
			case error:
				desc = rval.(error).Error()
			default:
				desc = fmt.Sprint(rval)
			}
			render.Render(rw, r, g.ErrRender(errors.New(desc)))

			func() {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("Report to sentry failed: %s, trace:\n%s", err, debug.Stack())
					}
				}()
				rec.monitor.errCounter.WithLabelValues(r.Method, r.URL.Path).Inc()
				switch rval := err.(type) {
				case error:
					rec.sentryClient.CaptureError(rval, nil)
				default:
					rec.sentryClient.CaptureError(errors.New(fmt.Sprint(rval)), nil)
				}
			}()
		}
	}()

	rec.h.ServeHTTP(rw, r)
}
