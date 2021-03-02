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

// create gateway service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_gateway_go.RegisterGatewayServer(s, NewServer())
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8001")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8001",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = tm2_proto_gateway_go.RegisterGatewayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8000",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8000")
	log.Fatalln(gwServer.ListenAndServe())
}
