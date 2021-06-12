// Package main implements a server for Greeter service.
package main

import (
	"context"
	"go-postgres/helloworld"
	"log"
	"net"

	pb "go-postgres/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	//pb.UnimplementedGreeterServer
	helloworld.GreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	log.Printf("Page: %v", in.GetPage())
	log.Printf("Size: %v", in.GetSize())
	return &pb.HelloReply{
		Message: "Hello: " + in.GetName(),
		Data:    in.GetPage() + in.GetSize(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
