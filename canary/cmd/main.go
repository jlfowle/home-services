package main

import (
	"net/http"
	"os"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jlfowle/home-services/canary/internal/instrumenting"
	"github.com/jlfowle/home-services/canary/internal/logging"
	"github.com/jlfowle/home-services/canary/internal/transport"
	"github.com/jlfowle/home-services/canary/pkg/canary"

	"github.com/go-kit/kit/log"
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

	var svc canary.CanaryService
	svc = canary.NewCanaryService()
	svc = logging.NewLoggingMiddleware(logger, svc)
	svc = instrumenting.NewInstrumentingMiddleware(requestCount, requestLatency, svc)

	getHandler := transport.MakeGetEndpointHandler(svc)

	http.Handle("/", getHandler)
	http.Handle("/metrics", promhttp.Handler())
	_ = logger.Log("msg", "HTTP", "addr", ":8080")
	_ = logger.Log("err", http.ListenAndServe(":8080", nil))
}
