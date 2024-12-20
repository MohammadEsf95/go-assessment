package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "service2/contract/proto"
)

type server struct {
	pb.UnimplementedGetDataFromService2Server
}

func (s *server) GetData(_ context.Context, request *pb.Service2Request) (*pb.Service2Response, error) {
	data := map[int64]string{
		1: "one",
		2: "two",
		3: "three",
	}

	resp, ok := data[request.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	return &pb.Service2Response{Message: resp}, nil
}

// TODO correct contract package
func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGetDataFromService2Server(s, &server{})
	log.Printf("cmd listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
