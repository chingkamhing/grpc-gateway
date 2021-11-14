package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

//
// Reference:
// - mTLS: https://github.com/islishude/grpc-mtls-example
//

const serverAddr = "0.0.0.0:9000"
const userAddr = "user:9000"
const companyAddr = "company:9000"
const caFile = "certs/localhost/ca.crt"
const crtFile = "certs/localhost/server.crt"
const keyFile = "certs/localhost/server.key"

var isGRPSSecure = env("GRPC_SECURE", "no")

// create gateway service
func main() {
	log.Printf("isGRPSSecure: %v", isGRPSSecure)
	serviceOptions := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// user gRPC client connection
	userConn, err := grpc.DialContext(context.Background(), userAddr, serviceOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// company gRPC client connection
	companyConn, err := grpc.DialContext(context.Background(), companyAddr, serviceOptions...)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	tlsCredentials, err := loadTLSCredentials(caFile, crtFile, keyFile)
	if err != nil {
		log.Fatalln("load TLS credentials error:", err)
	}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(middlewareLog),
	}
	if isGRPSSecure == "yes" {
		// Enable TLS for all incoming connections.
		serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	}
	server := grpc.NewServer(serverOptions...)
	// Attach the Greeter service to the server
	tm2_proto_gateway_go.RegisterGatewayServer(server, NewServer(userConn, companyConn))
	// Serve gRPC server
	log.Println("Serving gRPC on ", serverAddr)
	log.Fatalln(server.Serve(lis))
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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	}
	return credentials.NewTLS(config), nil
}

func middlewareLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// get client tls info
	if p, ok := peer.FromContext(ctx); ok {
		mtls, ok := p.AuthInfo.(credentials.TLSInfo)
		if ok {
			for _, item := range mtls.State.PeerCertificates {
				log.Println("request certificate subject:", item.Subject)
			}
		}
	}
	return handler(ctx, req)
}

func env(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
