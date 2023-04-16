package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/jlfowle/home-services/zilch/pkg/zilch"
)

type Set interface {
	GetEndpoint() endpoint.Endpoint
}

type set struct {
	getEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc zilch.ZilchService) Set {
	return &set{
		getEndpoint: makeGetEndpoint(svc),
	}
}

type getResponse struct{}

func makeGetEndpoint(svc zilch.ZilchService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return getResponse{}, svc.Get()
	}
}

func (s *set) GetEndpoint() endpoint.Endpoint {
	return s.getEndpoint
}
