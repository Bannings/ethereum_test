package middleware

import (
	"bytes"
	"net/http"
	"text/template"
	"time"

	"github.com/eddyzhou/log"
)

const (
	DefaultFormat     = "{{.StartTime}} | {{.Status}} | {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} "
	DefaultDateFormat = time.RFC3339
)

type options struct {
	logFormat  string
	dateFormat string
}

type LoggingOption func(*options)

func Format(logFormat, dateFormat string) LoggingOption {
	return func(options *options) {
		options.logFormat = logFormat
		options.dateFormat = dateFormat
	}
}

type LogEntry struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

type HttpLogger struct {
	*log.Logger
	opt      *options
	template *template.Template
}

func NewHttpLogger(logger *log.Logger, opts ...LoggingOption) *HttpLogger {
	opt := &options{}
	for _, o := range opts {
		o(opt)
	}

	if opt.dateFormat == "" {
		opt.dateFormat = DefaultDateFormat
	}

	lf := opt.logFormat
	if opt.logFormat == "" {
		lf = DefaultFormat
	}
	t := template.Must(template.New("r2_parser").Parse(lf))

	return &HttpLogger{
		Logger:   logger,
		opt:      opt,
		template: t,
	}
}

func Logging(logger *log.Logger, opts ...LoggingOption) func(next http.Handler) http.Handler {
	httpLogger := NewHttpLogger(logger, opts...)

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			rw := newCustomResponseWriter(w)
			start := time.Now()

			defer func() {
				log := LogEntry{
					StartTime: start.Format(httpLogger.opt.dateFormat),
					Status:    rw.status,
					Duration:  time.Since(start),
					Hostname:  r.Host,
					Method:    r.Method,
					Path:      r.URL.Path,
					Request:   r,
				}

				buff := &bytes.Buffer{}
				httpLogger.template.Execute(buff, log)
				httpLogger.Println(buff.String())
			}()

			next.ServeHTTP(rw, r)
		}
		return http.HandlerFunc(fn)
	}
}

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (c *customResponseWriter) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func newCustomResponseWriter(w http.ResponseWriter) *customResponseWriter {
	return &customResponseWriter{
		ResponseWriter: w,
		status:         200,
	}
}
