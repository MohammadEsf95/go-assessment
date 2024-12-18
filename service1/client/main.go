package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "service1/contract"
)

func main() {
	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewGetDataFromService1Client(conn)

	data, err := client.GetData(context.Background(), &pb.Service1Request{Id: 2})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("received data from service1: ", data.GetMessage())
}
