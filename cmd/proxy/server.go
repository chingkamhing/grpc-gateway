package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/gen/tm2-proto-company-go"
	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/gen/tm2-proto-gateway-go"
	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/gen/tm2-proto-user-go"
)

type server struct {
	userClient    tm2_proto_user_go.UserClient
	companyClient tm2_proto_company_go.CompanyClient

	// Embed the unimplemented server
	tm2_proto_gateway_go.UnimplementedGatewayServer
}

var _ tm2_proto_gateway_go.GatewayServer = &server{}

// NewServer return GatewayServer interface
func NewServer(userConn *grpc.ClientConn, companyConn *grpc.ClientConn) tm2_proto_gateway_go.GatewayServer {
	return &server{
		userClient:    tm2_proto_user_go.NewUserClient(userConn),
		companyClient: tm2_proto_company_go.NewCompanyClient(companyConn),
	}
}

func (s *server) Login(ctx context.Context, request *tm2_proto_gateway_go.LoginRequest) (*tm2_proto_gateway_go.LoginReply, error) {
	//FIXME
	return nil, fmt.Errorf("method not implemented")
}

func (s *server) Logout(ctx context.Context, request *tm2_proto_gateway_go.LogoutRequest) (*tm2_proto_gateway_go.LogoutReply, error) {
	//FIXME
	return nil, fmt.Errorf("method not implemented")
}

func (s *server) CreateUser(ctx context.Context, request *tm2_proto_gateway_go.CreateUserRequest) (*tm2_proto_gateway_go.CreateUserReply, error) {
	userReply, err := s.userClient.CreateUser(ctx, &tm2_proto_user_go.CreateUserRequest{Value: request.Value.User})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "CreateUser error: %v", errMsg.Err())
	}
	companyReply, err := s.companyClient.CreateCompany(ctx, &tm2_proto_company_go.CreateCompanyRequest{Value: request.Value.Company})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "CreateCompany error: %v", errMsg.Err())
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
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "ListUser error: %v", errMsg.Err())
	}
	companyReply, err := s.companyClient.ListCompany(ctx, &tm2_proto_company_go.ListCompanyRequest{Offset: request.Offset, Limit: request.Limit})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "ListCompany error: %v", errMsg.Err())
	}
	reply := &tm2_proto_gateway_go.ListUserReply{
		Values: []*tm2_proto_gateway_go.UserDetail{},
		Offset: userReply.Offset,
		Limit:  userReply.Limit,
		Count:  userReply.Count,
	}
	for i := range userReply.Values {
		reply.Values = append(reply.Values, &tm2_proto_gateway_go.UserDetail{
			User:    userReply.Values[i],
			Company: companyReply.Values[i],
		})
	}
	return reply, nil
}

func (s *server) GetUser(ctx context.Context, request *tm2_proto_gateway_go.GetUserRequest) (*tm2_proto_gateway_go.GetUserReply, error) {
	userReply, err := s.userClient.GetUser(ctx, &tm2_proto_user_go.GetUserRequest{Id: request.Id})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "GetUser error: %v", errMsg.Err())
	}
	companyReply, err := s.companyClient.GetCompany(ctx, &tm2_proto_company_go.GetCompanyRequest{Id: request.Id})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "GetCompany error: %v", errMsg.Err())
	}
	reply := &tm2_proto_gateway_go.GetUserReply{
		Value: &tm2_proto_gateway_go.UserDetail{
			User:    userReply.Value,
			Company: companyReply.Value,
		},
	}
	return reply, nil
}

func (s *server) UpdateUser(ctx context.Context, request *tm2_proto_gateway_go.UpdateUserRequest) (*tm2_proto_gateway_go.UpdateUserReply, error) {
	userReply, err := s.userClient.UpdateUser(ctx, &tm2_proto_user_go.UpdateUserRequest{Value: request.Value.User})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "UpdateUser error: %v", errMsg.Err())
	}
	companyReply, err := s.companyClient.UpdateCompany(ctx, &tm2_proto_company_go.UpdateCompanyRequest{Value: request.Value.Company})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "UpdateCompany error: %v", errMsg.Err())
	}
	reply := &tm2_proto_gateway_go.UpdateUserReply{
		Value: &tm2_proto_gateway_go.UserDetail{
			User:    userReply.Value,
			Company: companyReply.Value,
		},
	}
	return reply, nil
}

func (s *server) DeleteUser(ctx context.Context, request *tm2_proto_gateway_go.DeleteUserRequest) (*tm2_proto_gateway_go.DeleteUserReply, error) {
	userReply, err := s.userClient.DeleteUser(ctx, &tm2_proto_user_go.DeleteUserRequest{Id: request.Id})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "DeleteUser error: %v", errMsg.Err())
	}
	_, err = s.companyClient.DeleteCompany(ctx, &tm2_proto_company_go.DeleteCompanyRequest{Id: request.Id})
	if err != nil {
		errMsg, _ := status.FromError(err)
		return nil, status.Errorf(codes.NotFound, "DeleteCompany error: %v", errMsg.Err())
	}
	reply := &tm2_proto_gateway_go.DeleteUserReply{
		Id: userReply.Id,
	}
	return reply, nil
}
