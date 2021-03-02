package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_gateway_go.UnimplementedGatewayServer
}

// NewServer return GatewayServer interface
func NewServer() tm2_proto_gateway_go.GatewayServer {
	return &server{}
}

func (s *server) CreateUser(context.Context, *tm2_proto_gateway_go.CreateUserRequest) (*tm2_proto_gateway_go.CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (s *server) ListUser(context.Context, *tm2_proto_gateway_go.ListUserRequest) (*tm2_proto_gateway_go.ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (s *server) GetUser(context.Context, *tm2_proto_gateway_go.GetUserRequest) (*tm2_proto_gateway_go.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (s *server) UpdateUser(context.Context, *tm2_proto_gateway_go.UpdateUserRequest) (*tm2_proto_gateway_go.UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *server) DeleteUser(context.Context, *tm2_proto_gateway_go.DeleteUserRequest) (*tm2_proto_gateway_go.DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
