package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

func MakeGetEndpointHandler(svc zilch.ZilchService) httptransport.Server {
	return *httptransport.NewServer(
		makeGetEndpoint(svc),
		decodeGetRequest,
		encodeGetResponse,
	)
}

func makeGetEndpoint(svc zilch.ZilchService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return getResponse{}, svc.Get()
	}
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(&getResponse{})
}

type getResponse struct{}
