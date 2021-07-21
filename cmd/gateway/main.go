package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	tm2_proto_gateway_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-gateway-go"
)

var errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
var errInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")

const webHost = "0.0.0.0:8000"
const gatewayHost = "0.0.0.0:8001"
const userHost = "0.0.0.0:8002"
const companyHost = "0.0.0.0:8003"

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
	lis, err := net.Listen("tcp", gatewayHost)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	cert, err := tls.LoadX509KeyPair("deploy/cert/localhost/localhost.crt", "deploy/cert/localhost/localhost.key")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	s := grpc.NewServer(serverOptions...)
	// Attach the Greeter service to the server
	tm2_proto_gateway_go.RegisterGatewayServer(s, NewServer(userConn, companyConn))
	// Serve gRPC server
	log.Println("Serving gRPC on ", gatewayHost)
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

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
	gatewayConn, err := grpc.DialContext(context.Background(), gatewayHost, gatewayOptions...)
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
		Addr:    webHost,
		Handler: gwmux,
	}

	log.Printf("Serving gRPC-Gateway on http://%s\n", webHost)
	log.Fatalln(gwServer.ListenAndServe())
}

// valid validates the authorization.
func valid(authorization []string) bool {
	log.Printf("authorization: %v", authorization)
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	log.Printf("token: %v", token)
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "1234567890abcdefg"
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	log.Printf("md: %v", md)
	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
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
