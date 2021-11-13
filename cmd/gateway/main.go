package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

const host = "0.0.0.0"
const port = 8000
const proxyHost = "proxy"
const proxyPort = 9000

// create gateway service
func main() {
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	creds, err := credentials.NewClientTLSFromFile("deploy/cert/localhost/ca.pem", "localhost")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	gatewayOptions := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(authInterceptor),
		// oauth.NewOauthAccess requires the configuration of transport credentials.
		grpc.WithTransportCredentials(creds),
	}
	gatewayConn, err := grpc.DialContext(context.Background(), fmt.Sprintf("%s:%d", proxyHost, proxyPort), gatewayOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwMux := runtime.NewServeMux(
		// hand-pick to forward http method and url path to metadata in order to be used in authentication middleware
		runtime.WithMetadata(func(c context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("x-forwarded-method", req.Method)
		}),
		runtime.WithMetadata(func(c context.Context, req *http.Request) metadata.MD {
			return metadata.Pairs("x-forwarded-url-path", req.URL.Path)
		}),
	)
	// Register Greeter
	err = tm2_proto_gateway_go.RegisterGatewayHandler(context.Background(), gwMux, gatewayConn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	gwServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: gwMux,
	}
	log.Printf("Serving http gateway on http://%s:%d\n", host, port)
	log.Fatalln(gwServer.ListenAndServe())
}

// authInterceptor authenticate endpoint access
func authInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// get out bound metadata
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	log.Printf("md: %#v", md)
	// check if authorization is needed
	auths, ok := md["authorization"]
	if !ok {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
	// need authorization
	for _, auth := range auths {
		//FIXME, check authorization here
		log.Printf("auth: %#v", auth)
	}
	log.Printf("req: %#v", req)
	return invoker(ctx, method, req, reply, cc, opts...)
}
