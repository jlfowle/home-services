package transport

import (
	"context"

	"github.com/unnamed/home-services/canary/api/v1/canary"

	"github.com/unnamed/home-services/canary/pkg/canary/endpoints"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcServer struct {
	health grpctransport.Handler
	ping   grpctransport.Handler
	canary.UnimplementedCanaryServiceServer
}

func NewGRPCServer(ep endpoints.Set) canary.CanaryServiceServer {
	return &grpcServer{
		health: grpctransport.NewServer(
			ep.HealthEndpoint,
			decodeGRPCHealthRequest,
			decodeGRPCHealthResponse,
		),
		ping: grpctransport.NewServer(
			ep.PingEndpoint,
			decodeGRPCPingRequest,
			decodeGRPCPingResponse,
		),
	}
}

func (g *grpcServer) Health(ctx context.Context, r *emptypb.Empty) (*canary.HealthReply, error) {
	_, rep, err := g.health.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*canary.HealthReply), nil
}

func (g *grpcServer) Ping(ctx context.Context, r *emptypb.Empty) (*canary.PingReply, error) {
	_, rep, err := g.ping.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*canary.PingReply), nil
}

func decodeGRPCHealthRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoints.HealthRequest{}, nil
}

func decodeGRPCPingRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoints.PingRequest{}, nil
}

func decodeGRPCHealthResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.HealthResponse)
	return &canary.HealthReply{Code: int64(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCPingResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.PingResponse)
	return &canary.PingReply{Output: reply.Output}, nil
}
