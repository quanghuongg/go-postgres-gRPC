package main

import (
	"context"
	"fmt"
	pb "go-postgres/bidaspin"
	"go-postgres/bidaspin/gift"
	"go-postgres/bidaspin/redis"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	port = ":50050"
)

type BidaSpinServer struct {
	pb.BidaSpinServer
}

func (s *BidaSpinServer) UpdateTotalSpin(ctx context.Context, rq *pb.SpinRequest) (*pb.SpinResponse, error) {
	uid := rq.UserId
	count := int(rq.Count)
	strUid := strconv.Itoa(int(uid))
	rdb := redis.RedisClient()
	val, err := rdb.Get(ctx, strUid).Result()
	if err != nil {
	} else {
		intVar, _ := strconv.Atoi(val)
		count = count + intVar
		err := rdb.Set(ctx, strUid, strconv.Itoa(count), 0).Err()
		if err != nil {
			panic(err)
		}
	}
	return &pb.SpinResponse{
		Message: "success",
		Data:    strconv.Itoa(count),
	}, nil
}
func (s *BidaSpinServer) GetTotalSpin(ctx context.Context, rq *pb.SpinRequest) (*pb.SpinResponse, error) {
	uid := rq.UserId
	strUid := strconv.Itoa(int(uid))
	rdb := redis.RedisClient()
	val, err := rdb.Get(ctx, strUid).Result()
	if err != nil {
		val = "0"
		err := rdb.Set(ctx, strUid, "0", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	return &pb.SpinResponse{
		Message: "success",
		Data:    val,
	}, nil
}

func (s *BidaSpinServer) DoSpin(ctx context.Context, rq *pb.SpinRequest) (*pb.SpinResponse, error) {
	uid := rq.UserId
	strUid := strconv.Itoa(int(uid))
	rdb := redis.RedisClient()
	val, err := rdb.Get(ctx, strUid).Result()
	var data = "Total spin invalid"
	var msg = "Fail"
	if err == nil {
		intVar, _ := strconv.Atoi(val)
		if intVar > 0 {
			randomGift := gift.RandomGift()
			data = randomGift.GiftToJsonString()
			msg = "success"
			//update
			err := rdb.Set(ctx, strUid, strconv.Itoa(intVar-1), 0).Err()
			if err != nil {
				panic(err)
			} else {
				fmt.Println("Update total spin success")
			}
		}
	}
	return &pb.SpinResponse{
		Message: msg,
		Data:    data,
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBidaSpinServer(s, &BidaSpinServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Println("BidaH5Spin server ready")
	}
}
