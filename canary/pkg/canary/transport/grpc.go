package transport

import (
	"context"

	"github.com/jlfowle/home-services/canary/api/v1/canary"

	"github.com/jlfowle/home-services/canary/pkg/canary/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	get           grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) canary.CanaryServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *canary.GetRequest) (*canary.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*canary.GetReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *canary.ServiceStatusRequest) (*canary.ServiceStatusReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*canary.ServiceStatusReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.GetRequest{}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	_ = grpcReply.(*canary.GetReply)
	return endpoints.GetResponse{}, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*canary.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
