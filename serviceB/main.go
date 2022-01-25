package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	protoA "learn-service-communication/proto-repo/protoA"
	"log"
	"net"
)

type server struct {
	name string
}

func MessageService() *server {
	return &server{
		name: "Server",
	}
}

var (
	port = flag.String("port", "8000", "port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("ServiceB listening in: ", *port)
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	serviceB := MessageService()
	protoA.RegisterMessageServiceServer(grpcServer, serviceB)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Get(c context.Context, req *protoA.Request) (*protoA.Response, error) {
	log.Println("incoming req : ", req.Name)
	response := &protoA.Response{
		Message: fmt.Sprintf("Hi %s! Welcome Back", req.Name),
	}

	return response, nil
}
