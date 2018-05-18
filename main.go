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

	"gitlab.chainedfinance.com/chaincore/r2/g"
	"gitlab.chainedfinance.com/chaincore/r2/handler"
	mw "gitlab.chainedfinance.com/chaincore/r2/middleware"

	"github.com/eddyzhou/log"
	"github.com/getsentry/raven-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	sentryDSN = "http://91f3f829417f4780970565b9916345be:033365bfe6484634b3539b8e0bb22019@sentry.chainedfinance.com/3"
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

	r.Route("/api/cf/ar", func(r chi.Router) {
		r.Use(mw.ApiVersionCtx("v1"))
		r.Use(mw.Auth())

		r.Get("/queryCompanyByCId", handler.QueryCompanyHandler)
		r.Post("/addCompany", handler.AddCompanyHandler)
		r.Post("/updateCompanyInfo", handler.UpdateCompanyHander)

		r.Get("/queryARByArid", handler.QueryArHandler)
		r.Post("/createAR", handler.CreateArHandler)
		r.Post("/updateARInfo", handler.UpdateArHander)

		r.Get("/queryContractByContractNo", handler.QueryContractHandler)
		r.Post("/addContract", handler.AddContractHandler)
		r.Post("/updateContractInfo", handler.UpdateContractHander)

		r.Get("/queryARPaymentTx", handler.QueryArPayHandler)
		r.Post("/payByAR", handler.CreateArPayHandler)
		r.Post("/updatePaymentTxOfAR", handler.UpdateArPayHander)

		r.Get("/queryARDiscountTx", handler.QueryDiscountHandler)
		r.Post("/discountByAR", handler.CreateDiscountHandler)
		r.Post("/updateDiscountTxOfAR", handler.UpdateDiscountHander)
	})

	r.Route("/api/cf/node", func(r chi.Router) {
		r.Use(mw.Auth())
		r.Use(mw.OnlyCF)

		r.Post("/deleteNode/", handler.DeleteNodeHandler)
		r.Post("/addNode/", handler.AddNodeHandler)
		r.Get("/getNodeKey/", handler.QueryNodeHandler)
	})

	r.Route("/api/fx", func(r chi.Router) {
		r.Use(mw.Auth())
		r.Use(mw.OnlyCF)

		r.Post("/supplier/register", handler.RegisterSupplierHandler)
	})

	r.Mount("/debug", middleware.Profiler())

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

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
