package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	tm2_proto_user_go "github.com/chingkamhing/grpc-gateway/gen/srv-proto-user-go"
)

const host = "0.0.0.0:9000"

// create user service
func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	tm2_proto_user_go.RegisterUserServer(s, NewServer())
	// Serve gRPC Server
	log.Println("Serving gRPC on ", host)
	log.Fatal(s.Serve(lis))
}
