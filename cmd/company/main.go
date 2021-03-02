package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	tm2_proto_company_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-company-go"
)

type server struct {
	// Embed the unimplemented server
	tm2_proto_company_go.UnimplementedCompanyServer
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

// create user service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_company_go.RegisterCompanyServer(s, &server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8003")
	log.Fatal(s.Serve(lis))
}
