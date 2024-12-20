package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "service1/contract/proto"
)

type server struct {
	pb.UnimplementedGetDataFromService1Server
}

func (s *server) GetData(_ context.Context, request *pb.Service1Request) (*pb.Service1Response, error) {
	fmt.Println("server 1 inja")
	data := map[int64]string{
		1: "one",
		2: "two",
		3: "three",
	}

	resp, ok := data[request.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found")
	}
	return &pb.Service1Response{Message: resp}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGetDataFromService1Server(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
