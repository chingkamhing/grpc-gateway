package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

const host = "0.0.0.0:9000"
const userHost = "user:9000"
const companyHost = "company:9000"
const certFile = "deploy/cert/localhost/localhost.crt"
const keyFile = "deploy/cert/localhost/localhost.key"

// create gateway service
func main() {
	serviceOptions := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// user gRPC client connection
	userConn, err := grpc.DialContext(context.Background(), userHost, serviceOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// company gRPC client connection
	companyConn, err := grpc.DialContext(context.Background(), companyHost, serviceOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	// cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	// if err != nil {
	// 	log.Fatalf("failed to load key pair: %s", err)
	// }
	serverOptions := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		// grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(serverOptions...)
	// Attach the Greeter service to the server
	tm2_proto_gateway_go.RegisterGatewayServer(s, NewServer(userConn, companyConn))
	// Serve gRPC server
	log.Println("Serving gRPC on ", host)
	log.Fatalln(s.Serve(lis))
}
