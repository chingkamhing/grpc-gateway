package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/lib/tm2-proto-user-go"
)

const userHost = "0.0.0.0:8002"

// create user service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", userHost)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_user_go.RegisterUserServer(s, NewServer())
	// Serve gRPC Server
	log.Println("Serving gRPC on ", userHost)
	log.Fatal(s.Serve(lis))
}
