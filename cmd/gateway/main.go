package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

const webHost = "0.0.0.0:8000"
const gatewayHost = "0.0.0.0:8001"
const userHost = "0.0.0.0:8002"
const companyHost = "0.0.0.0:8003"

// create gateway service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", gatewayHost)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	gatewayConn, err := grpc.DialContext(
		context.Background(),
		gatewayHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// user gRPC client connection
	userConn, err := grpc.DialContext(
		context.Background(),
		userHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// company gRPC client connection
	companyConn, err := grpc.DialContext(
		context.Background(),
		companyHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_gateway_go.RegisterGatewayServer(s, NewServer(userConn, companyConn))
	// Serve gRPC server
	log.Println("Serving gRPC on ", gatewayHost)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = tm2_proto_gateway_go.RegisterGatewayHandler(context.Background(), gwmux, gatewayConn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    webHost,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://%s\n", webHost)
	log.Fatalln(gwServer.ListenAndServe())
}
