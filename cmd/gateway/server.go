package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-company-go"
	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-user-go"
)

type server struct {
	userClient    tm2_proto_user_go.UserClient
	companyClient tm2_proto_company_go.CompanyClient

	// Embed the unimplemented server
	tm2_proto_gateway_go.UnimplementedGatewayServer
}

// NewServer return GatewayServer interface
func NewServer(userConn *grpc.ClientConn, companyConn *grpc.ClientConn) tm2_proto_gateway_go.GatewayServer {
	return &server{
		userClient:    tm2_proto_user_go.NewUserClient(userConn),
		companyClient: tm2_proto_company_go.NewCompanyClient(companyConn),
	}
}

func (s *server) CreateUser(ctx context.Context, request *tm2_proto_gateway_go.CreateUserRequest) (*tm2_proto_gateway_go.CreateUserReply, error) {
	userReply, err := s.userClient.CreateUser(ctx, &tm2_proto_user_go.CreateUserRequest{Value: request.Value.User})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "CreateUser error")
	}
	companyReply, err := s.companyClient.CreateCompany(ctx, &tm2_proto_company_go.CreateCompanyRequest{Value: request.Value.Company})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "CreateCompany error")
	}
	reply := &tm2_proto_gateway_go.CreateUserReply{
		Value: &tm2_proto_gateway_go.UserDetail{
			User:    userReply.Value,
			Company: companyReply.Value,
		},
	}
	return reply, nil
}
func (s *server) ListUser(ctx context.Context, request *tm2_proto_gateway_go.ListUserRequest) (*tm2_proto_gateway_go.ListUserReply, error) {
	userReply, err := s.userClient.ListUser(ctx, &tm2_proto_user_go.ListUserRequest{Offset: request.Offset, Limit: request.Limit})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "ListUser error: %v", err)
	}
	companyReply, err := s.companyClient.ListCompany(ctx, &tm2_proto_company_go.ListCompanyRequest{Offset: request.Offset, Limit: request.Limit})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "ListCompany error: %v", err)
	}
	reply := &tm2_proto_gateway_go.ListUserReply{
		Value:  []*tm2_proto_gateway_go.UserDetail{},
		Offset: userReply.Offset,
		Limit:  userReply.Limit,
		Count:  userReply.Count,
	}
	for i := range userReply.Value {
		reply.Value = append(reply.Value, &tm2_proto_gateway_go.UserDetail{
			User:    userReply.Value[i],
			Company: companyReply.Value[i],
		})
	}
	return reply, nil
}
func (s *server) GetUser(ctx context.Context, request *tm2_proto_gateway_go.GetUserRequest) (*tm2_proto_gateway_go.GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (s *server) UpdateUser(ctx context.Context, request *tm2_proto_gateway_go.UpdateUserRequest) (*tm2_proto_gateway_go.UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *server) DeleteUser(ctx context.Context, request *tm2_proto_gateway_go.DeleteUserRequest) (*tm2_proto_gateway_go.DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
