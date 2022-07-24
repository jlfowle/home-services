package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jlfowle/home-services/canary/pkg/canary"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeGetEndpointHandler(svc canary.CanaryService) httptransport.Server {
	return *httptransport.NewServer(
		makeGetServiceHealth(svc),
		decodeGetServiceHealthRequest,
		encodeGetServiceHealthResponse,
	)
}

func makeGetServiceHealth(svc canary.CanaryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return svc.GetServiceHealth()
	}
}

func decodeGetServiceHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetServiceHealthResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(&getServiceHealthResponse{})
}

type getServiceHealthResponse struct{}
