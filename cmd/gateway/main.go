package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/gen/tm2-proto-gateway-go"
)

//
// Reference:
// - mTLS: https://github.com/islishude/grpc-mtls-example
//

const serverAddr = "0.0.0.0:8000"
const proxyAddr = "proxy:9000"
const caFile = "certs/localhost/ca.crt"
const crtFile = "certs/localhost/client.crt"
const keyFile = "certs/localhost/client.key"

var isGRPCSecure = env("GRPC_SECURE", "no")

// create gateway service
func main() {
	log.Printf("isGRPCSecure: %v", isGRPCSecure)
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	tlsCredentials, err := loadTLSCredentials(caFile, crtFile, keyFile)
	if err != nil {
		log.Fatalln("load TLS credentials error:", err)
	}
	gatewayOptions := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(authInterceptor),
	}
	if isGRPCSecure == "yes" {
		// oauth.NewOauthAccess requires the configuration of transport credentials.
		gatewayOptions = append(gatewayOptions, grpc.WithTransportCredentials(tlsCredentials))
	} else {
		gatewayOptions = append(gatewayOptions, grpc.WithInsecure())
	}
	gatewayConn, err := grpc.DialContext(context.Background(), proxyAddr, gatewayOptions...)
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
		Addr:    serverAddr,
		Handler: gwMux,
	}
	log.Printf("Serving http gateway on http://%s\n", serverAddr)
	log.Fatalln(gwServer.ListenAndServe())
}

func loadTLSCredentials(caFile, crtFile, keyFile string) (credentials.TransportCredentials, error) {
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("load certification error: %w", err)
	}
	// Load certificate of the CA who signed server's certificate
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	}
	return credentials.NewTLS(config), nil
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

func env(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
