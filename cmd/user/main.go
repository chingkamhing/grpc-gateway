package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-user-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_user_go.UnimplementedUserServer
}

func (s *server) CreateUser(context.Context, *tm2_proto_user_go.CreateUserRequest) (*tm2_proto_user_go.CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (s *server) ListUser(context.Context, *tm2_proto_user_go.ListUserRequest) (*tm2_proto_user_go.ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (s *server) GetUser(context.Context, *tm2_proto_user_go.GetUserRequest) (*tm2_proto_user_go.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (s *server) UpdateUser(context.Context, *tm2_proto_user_go.UpdateUserRequest) (*tm2_proto_user_go.UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *server) DeleteUser(context.Context, *tm2_proto_user_go.DeleteUserRequest) (*tm2_proto_user_go.DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

// create user service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_user_go.RegisterUserServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8001")
	log.Fatal(s.Serve(lis))
}
