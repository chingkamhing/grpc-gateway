package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-company-go"
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
	fmt.Printf("CreateCompany: %+v\n", request.Value)
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
	reply := &tm2_proto_company_go.GetCompanyReply{
		Value: companyMap[request.Id],
	}
	return reply, nil
}

func (s *server) UpdateCompany(ctx context.Context, request *tm2_proto_company_go.UpdateCompanyRequest) (*tm2_proto_company_go.UpdateCompanyReply, error) {
	if _, ok := companyMap[request.Value.CompanyID]; ok != true {
		return nil, status.Errorf(codes.NotFound, "UpdateCompany cannot find id: ", request.Value.CompanyID)
	}
	companyMap[request.Value.CompanyID] = request.Value
	reply := &tm2_proto_company_go.UpdateCompanyReply{
		Value: companyMap[request.Value.CompanyID],
	}
	return reply, nil
}

func (s *server) DeleteCompany(ctx context.Context, request *tm2_proto_company_go.DeleteCompanyRequest) (*tm2_proto_company_go.DeleteCompanyReply, error) {
	if _, ok := companyMap[request.Id]; ok != true {
		return nil, status.Errorf(codes.NotFound, "DeleteCompany cannot find id: ", request.Id)
	}
	delete(companyMap, request.Id)
	return &tm2_proto_company_go.DeleteCompanyReply{}, nil
}
