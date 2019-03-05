package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"gitlab.com/ethereum_test/g"
	"gitlab.com/ethereum_test/handler"
	mw "gitlab.com/ethereum_test/middleware"

	"github.com/eddyzhou/log"
	"github.com/getsentry/raven-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/ethereum_test/tokens"
)

const (
	sentryDSN = "http://f283fee151d848b8a21074d95a509472:3a35e8635ebb4351847ee7ba557eec39@193.112.186.211/2"
)

var (
	logLevel = flag.String("L", "info", "log level: info, debug, warn, error, fatal")
	logFile  = flag.String("logfile", "", "log file path")
	port     = flag.Int("port", 10088, "listen port")
)

func init() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	if len(*logFile) > 0 {
		log.Std = log.NewRotate(*logFile, "", log.Ldefault, lvl)
	} else {
		log.Std = log.New(os.Stderr, "", log.Ldefault, lvl)
	}

	if err := raven.SetDSN(sentryDSN); err != nil {
		log.Fatal(err)
	}

	configFile := args[0]
	if _, err := g.LoadConfig(configFile); err != nil {
		log.Fatal(err)
	}

	pid2file()
}

func pid2file() {
	file, err := os.OpenFile("r2.pid", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("open file err: %v", err)
	}
	defer file.Close()
	file.Write([]byte(strconv.Itoa(syscall.Getpid())))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [configFile]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	m := mw.NewMonitor("r2", *port)

	r := chi.NewRouter()
	r.Use(mw.Recovery(m))
	r.Use(mw.Monitoring(m))
	r.Use(mw.Logging(log.Std))

	h := promhttp.Handler()
	r.Get("/metrics", func(rw http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(rw, req)
	})

	r.Route("/api/token", func(r chi.Router) {

		r.Post("/supplier/register", handler.RegisterSupplierHandler)
		r.Post("/asset", handler.AssetHandler)
		r.Get("/query", handler.QueryTokenHandler)
	})

	r.Mount("/debug", middleware.Profiler())

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	done := make(chan struct{})
	defer close(done)
	go tokens.ExcuteTransaction(done)

	log.Infof("listening on: %v", *port)
	srv := &http.Server{Addr: fmt.Sprintf(":%d", *port), Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
