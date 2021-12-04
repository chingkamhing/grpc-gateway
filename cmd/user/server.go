package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/gen/srv-proto-user-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_user_go.UnimplementedUserServer
}

// save users in memory
var userMap = map[string]*tm2_proto_user_go.UserInfo{}

// NewServer return GatewayServer interface
func NewServer() tm2_proto_user_go.UserServer {
	return &server{}
}

func (s *server) CreateUser(ctx context.Context, request *tm2_proto_user_go.CreateUserRequest) (*tm2_proto_user_go.CreateUserReply, error) {
	userMap[request.Value.UserID] = request.Value
	reply := &tm2_proto_user_go.CreateUserReply{
		Value: userMap[request.Value.UserID],
	}
	return reply, nil
}

func (s *server) ListUser(ctx context.Context, request *tm2_proto_user_go.ListUserRequest) (*tm2_proto_user_go.ListUserReply, error) {
	users := []*tm2_proto_user_go.UserInfo{}
	for i := range userMap {
		users = append(users, userMap[i])
	}
	reply := &tm2_proto_user_go.ListUserReply{
		Values: users,
		Offset: request.Offset,
		Limit:  request.Limit,
		Count:  int32(len(users)),
	}
	return reply, nil
}

func (s *server) GetUser(ctx context.Context, request *tm2_proto_user_go.GetUserRequest) (*tm2_proto_user_go.GetUserReply, error) {
	id := request.Id
	value, ok := userMap[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %v not found", id)
	}
	reply := &tm2_proto_user_go.GetUserReply{
		Value: value,
	}
	return reply, nil
}

func (s *server) UpdateUser(ctx context.Context, request *tm2_proto_user_go.UpdateUserRequest) (*tm2_proto_user_go.UpdateUserReply, error) {
	id := request.Value.UserID
	if _, ok := userMap[id]; !ok {
		return nil, status.Errorf(codes.NotFound, "user %v not found: ", id)
	}
	userMap[id] = request.Value
	reply := &tm2_proto_user_go.UpdateUserReply{
		Value: userMap[id],
	}
	return reply, nil
}

func (s *server) DeleteUser(ctx context.Context, request *tm2_proto_user_go.DeleteUserRequest) (*tm2_proto_user_go.DeleteUserReply, error) {
	id := request.Id
	_, ok := userMap[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %v not found", id)
	}
	delete(userMap, request.Id)
	return &tm2_proto_user_go.DeleteUserReply{
		Id: id,
	}, nil
}
