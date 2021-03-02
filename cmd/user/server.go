package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-user-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_user_go.UnimplementedUserServer
}

var users = []*tm2_proto_user_go.UserInfo{}

// NewServer return GatewayServer interface
func NewServer() tm2_proto_user_go.UserServer {
	return &server{}
}

func (s *server) CreateUser(ctx context.Context, request *tm2_proto_user_go.CreateUserRequest) (*tm2_proto_user_go.CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (s *server) ListUser(ctx context.Context, request *tm2_proto_user_go.ListUserRequest) (*tm2_proto_user_go.ListUserReply, error) {
	reply := &tm2_proto_user_go.ListUserReply{
		Value:  users,
		Offset: request.Offset,
		Limit:  request.Limit,
		Count:  int32(len(users)),
	}
	return reply, nil
}

func (s *server) GetUser(ctx context.Context, request *tm2_proto_user_go.GetUserRequest) (*tm2_proto_user_go.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (s *server) UpdateUser(ctx context.Context, request *tm2_proto_user_go.UpdateUserRequest) (*tm2_proto_user_go.UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (s *server) DeleteUser(ctx context.Context, request *tm2_proto_user_go.DeleteUserRequest) (*tm2_proto_user_go.DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
