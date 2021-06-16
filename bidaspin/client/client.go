package main

import (
	"context"
	pb "go-postgres/bidaspin"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50050"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBidaSpinClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTotalSpin(ctx, &pb.SpinRequest{
		UserId: 1,
		Type:   "vip",
	})
	if err != nil {
		log.Fatalf("Could not GetTotalSpin: %v", err)
	}
	log.Printf("GetTotalSpin 1: %s", r)

	for i := 0; i < 20; i++ {
		r3, err := c.DoSpin(ctx, &pb.SpinRequest{
			UserId: 1,
		})
		if err != nil {
			log.Fatalf("Could not GetTotalSpin: %v", err)
		}
		log.Printf("DoSpin: %s", r3)
	}

	r2, err := c.GetTotalSpin(ctx, &pb.SpinRequest{
		UserId: 1,
	})
	if err != nil {
		log.Fatalf("Could not GetTotalSpin: %v", err)
	}
	log.Printf("GetTotalSpin 2: %s", r2)
}
