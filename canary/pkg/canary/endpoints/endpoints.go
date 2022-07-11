package endpoints

import (
	"context"
	"errors"
	"os"

	"github.com/jlfowle/home-services/canary/pkg/canary"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	CanaryEndpoint        endpoint.Endpoint
}

func NewEndpointSet(svc canary.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeGetEndpoint(svc canary.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := svc.Get(ctx)
		if err != nil {
			return GetResponse{err.Error()}, nil
		}
		return GetResponse{""}, nil
	}
}

func MakeServiceStatusEndpoint(svc canary.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context) error {
	resp, err := s.GetEndpoint(ctx, GetRequest{})
	if err != nil {
		return err
	}
	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return errors.New(getResp.Err)
	}
	return nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
