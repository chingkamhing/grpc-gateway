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

// NewServer return GatewayServer interface
func NewServer() tm2_proto_company_go.CompanyServer {
	return &server{}
}

func (s *server) CreateCompany(context.Context, *tm2_proto_company_go.CreateCompanyRequest) (*tm2_proto_company_go.CreateCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (s *server) ListCompany(context.Context, *tm2_proto_company_go.ListCompanyRequest) (*tm2_proto_company_go.ListCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompany not implemented")
}
func (s *server) GetCompany(context.Context, *tm2_proto_company_go.GetCompanyRequest) (*tm2_proto_company_go.GetCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (s *server) UpdateCompany(context.Context, *tm2_proto_company_go.UpdateCompanyRequest) (*tm2_proto_company_go.UpdateCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (s *server) DeleteCompany(context.Context, *tm2_proto_company_go.DeleteCompanyRequest) (*tm2_proto_company_go.DeleteCompanyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
