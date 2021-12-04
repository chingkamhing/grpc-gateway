package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/gen/tm2-proto-company-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_company_go.UnimplementedCompanyServer
}

// save companies in memory
var companyMap = map[string]*tm2_proto_company_go.CompanyInfo{}

// NewServer return GatewayServer interface
func NewServer() tm2_proto_company_go.CompanyServer {
	return &server{}
}

func (s *server) CreateCompany(ctx context.Context, request *tm2_proto_company_go.CreateCompanyRequest) (*tm2_proto_company_go.CreateCompanyReply, error) {
	companyMap[request.Value.CompanyID] = request.Value
	reply := &tm2_proto_company_go.CreateCompanyReply{
		Value: companyMap[request.Value.CompanyID],
	}
	return reply, nil
}

func (s *server) ListCompany(ctx context.Context, request *tm2_proto_company_go.ListCompanyRequest) (*tm2_proto_company_go.ListCompanyReply, error) {
	companys := []*tm2_proto_company_go.CompanyInfo{}
	for i := range companyMap {
		companys = append(companys, companyMap[i])
	}
	reply := &tm2_proto_company_go.ListCompanyReply{
		Values: companys,
		Offset: request.Offset,
		Limit:  request.Limit,
		Count:  int32(len(companys)),
	}
	return reply, nil
}

func (s *server) GetCompany(ctx context.Context, request *tm2_proto_company_go.GetCompanyRequest) (*tm2_proto_company_go.GetCompanyReply, error) {
	id := request.Id
	value, ok := companyMap[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "company %v not found", id)
	}
	reply := &tm2_proto_company_go.GetCompanyReply{
		Value: value,
	}
	return reply, nil
}

func (s *server) UpdateCompany(ctx context.Context, request *tm2_proto_company_go.UpdateCompanyRequest) (*tm2_proto_company_go.UpdateCompanyReply, error) {
	id := request.Value.CompanyID
	if _, ok := companyMap[id]; !ok {
		return nil, status.Errorf(codes.NotFound, "company %v not found: ", id)
	}
	companyMap[id] = request.Value
	reply := &tm2_proto_company_go.UpdateCompanyReply{
		Value: companyMap[id],
	}
	return reply, nil
}

func (s *server) DeleteCompany(ctx context.Context, request *tm2_proto_company_go.DeleteCompanyRequest) (*tm2_proto_company_go.DeleteCompanyReply, error) {
	id := request.Id
	_, ok := companyMap[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "company %v not found", id)
	}
	delete(companyMap, request.Id)
	return &tm2_proto_company_go.DeleteCompanyReply{}, nil
}
