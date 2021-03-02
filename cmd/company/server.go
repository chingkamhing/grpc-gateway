package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-company-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_company_go.UnimplementedCompanyServer
}

var companies = []*tm2_proto_company_go.CompanyInfo{}

// NewServer return GatewayServer interface
func NewServer() tm2_proto_company_go.CompanyServer {
	return &server{}
}

func (s *server) CreateCompany(ctx context.Context, request *tm2_proto_company_go.CreateCompanyRequest) (*tm2_proto_company_go.CreateCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}

func (s *server) ListCompany(ctx context.Context, request *tm2_proto_company_go.ListCompanyRequest) (*tm2_proto_company_go.ListCompanyReply, error) {
	reply := &tm2_proto_company_go.ListCompanyReply{
		Value:  companies,
		Offset: request.Offset,
		Limit:  request.Limit,
		Count:  int32(len(companies)),
	}
	return reply, nil
}

func (s *server) GetCompany(ctx context.Context, request *tm2_proto_company_go.GetCompanyRequest) (*tm2_proto_company_go.GetCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}

func (s *server) UpdateCompany(ctx context.Context, request *tm2_proto_company_go.UpdateCompanyRequest) (*tm2_proto_company_go.UpdateCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}

func (s *server) DeleteCompany(ctx context.Context, request *tm2_proto_company_go.DeleteCompanyRequest) (*tm2_proto_company_go.DeleteCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
