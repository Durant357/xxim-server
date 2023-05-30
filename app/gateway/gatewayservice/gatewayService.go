// Code generated by goctl. DO NOT EDIT!
// Source: gateway.proto

package gatewayservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GatewayApiRequest            = pb.GatewayApiRequest
	GatewayApiResponse           = pb.GatewayApiResponse
	GatewayGetUserConnectionReq  = pb.GatewayGetUserConnectionReq
	GatewayGetUserConnectionResp = pb.GatewayGetUserConnectionResp

	GatewayService interface {
		GatewayGetUserConnection(ctx context.Context, in *GatewayGetUserConnectionReq, opts ...grpc.CallOption) (*GatewayGetUserConnectionResp, error)
	}

	defaultGatewayService struct {
		cli zrpc.Client
	}
)

func NewGatewayService(cli zrpc.Client) GatewayService {
	return &defaultGatewayService{
		cli: cli,
	}
}

func (m *defaultGatewayService) GatewayGetUserConnection(ctx context.Context, in *GatewayGetUserConnectionReq, opts ...grpc.CallOption) (*GatewayGetUserConnectionResp, error) {
	client := pb.NewGatewayServiceClient(m.cli.Conn())
	return client.GatewayGetUserConnection(ctx, in, opts...)
}
