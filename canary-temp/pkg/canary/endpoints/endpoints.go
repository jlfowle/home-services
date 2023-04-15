package endpoints

import (
	"context"
	"errors"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/unnamed/home-services/canary/pkg/canary"
)

type Set struct {
	PingEndpoint   endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc canary.Service) Set {
	return Set{
		PingEndpoint:   MakePingEndpoint(svc),
		HealthEndpoint: MakeHealthEndpoint(svc),
	}
}

func MakePingEndpoint(svc canary.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		output := svc.Ping(ctx)
		return PingResponse{output}, nil
	}
}

func MakeHealthEndpoint(svc canary.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		code, err := svc.Health(ctx)
		if err != nil {
			return HealthResponse{Code: code, Err: err.Error()}, nil
		}
		return HealthResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Ping(ctx context.Context) string {
	resp, _ := s.PingEndpoint(ctx, PingRequest{})
	pingResp := resp.(PingResponse)
	return pingResp.Output
}

func (s *Set) Health(ctx context.Context) (int, error) {
	resp, err := s.HealthEndpoint(ctx, HealthRequest{})
	svcHealthResp := resp.(HealthResponse)
	if err != nil {
		return svcHealthResp.Code, err
	}
	if svcHealthResp.Err != "" {
		return svcHealthResp.Code, errors.New(svcHealthResp.Err)
	}
	return svcHealthResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
