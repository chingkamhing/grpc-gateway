package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

const host = "0.0.0.0:8000"
const proxy = "proxy:9000"

// create gateway service
func main() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	creds, err := credentials.NewClientTLSFromFile("deploy/cert/localhost/ca.pem", "localhost")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	gatewayOptions := []grpc.DialOption{
		// oauth.NewOauthAccess requires the configuration of transport credentials.
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(tokenAuth{token: "1234567890abcdefg"}),
	}
	gatewayConn, err := grpc.DialContext(context.Background(), proxy, gatewayOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = tm2_proto_gateway_go.RegisterGatewayHandler(context.Background(), gwmux, gatewayConn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    host,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://%s\n", host)
	log.Fatalln(gwServer.ListenAndServe())
}

type tokenAuth struct {
	token string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}
