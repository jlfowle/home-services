package main

import (
	"net/http"
	"os"

	"github.com/jlfowle/home-services/zilch/internal/endpoints"
	"github.com/jlfowle/home-services/zilch/internal/instrumenting"
	"github.com/jlfowle/home-services/zilch/internal/logging"
	"github.com/jlfowle/home-services/zilch/internal/transport"
	"github.com/jlfowle/home-services/zilch/pkg/zilch"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	fieldKeys := []string{}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "http",
		Subsystem: "request",
		Name:      "total",
		Help:      "Number of requests recieved.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "http",
		Subsystem: "request",
		Name:      "duration_seconds",
		Help:      "Total duration of requests in seconds.",
	}, fieldKeys)

	var svc zilch.ZilchService
	svc = zilch.NewZilchService()
	svc = logging.NewLoggingMiddleware(logger, svc)
	svc = instrumenting.NewInstrumentingMiddleware(requestCount, requestLatency, svc)

	eps := endpoints.NewEndpointSet(svc)
	getHandler := transport.MakeGetEndpointHandler(eps)

	http.Handle("/", getHandler)
	http.Handle("/metrics", promhttp.Handler())
	_ = logger.Log("msg", "HTTP", "addr", ":8080")
	_ = logger.Log("err", http.ListenAndServe(":8080", nil))
}
