package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/unnamed/home-services/canary/pkg/canary/endpoints"
)

func NewHttpHandler(ep endpoints.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healthz", httptransport.NewServer(
		ep.HealthEndpoint,
		decodeHTTPHealthRequest,
		encodeResponse,
	))
	m.Handle("/ping", httptransport.NewServer(
		ep.PingEndpoint,
		decodeHTTPPingRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPPingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.PingRequest
	if r.ContentLength == 0 {
		_ = logger.Log("msg", "Get request with no body")
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPHealthRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.HealthRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
