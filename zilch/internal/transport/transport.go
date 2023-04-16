package transport

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/jlfowle/home-services/zilch/internal/endpoints"
)

func MakeGetEndpointHandler(eps endpoints.Set) httptransport.Server {
	return *httptransport.NewServer(
		eps.GetEndpoint(),
		decodeGetRequest,
		encodeGetResponse,
	)
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeGetResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(&getResponse{})
}

type getResponse struct{}
