package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	contract "service3/contract/proto"
	"service3/entity"
	"service3/infrastructure/database"
	"service3/repository"
	"sync"
	"time"
)

// Sample data in order to insert to database
var sampleData = []entity.Model{
	{
		Name:      "NameOne",
		CreatedAt: time.Now(),
	},
	{
		Name:      "NameTwo",
		CreatedAt: time.Now(),
	},
	{
		Name:      "NameThree",
		CreatedAt: time.Now(),
	},
}

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Do the migrations and create tables
	err = db.AutoMigrate(&entity.Model{})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)

	connSrv1, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer connSrv1.Close()

	ctx := context.Background()
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)
	resultChan := make(chan *contract.Service1Response)

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting serviceOne goroutine")
		serviceOneResp, err1 := getDataFromServiceOneClient(ctx, connSrv1)
		if err1 != nil {
			errorChan <- fmt.Errorf("getDataFromServiceOneClient error: %w", err1)
			return
		}
		resultChan <- serviceOneResp
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = databaseOperation(repo); err != nil {
			errorChan <- fmt.Errorf("databaseOperation error: %w", err)
		}
		log.Println("db goroutine finished")
	}()

	wg.Wait()
	close(resultChan)
	close(errorChan)

	for err = range errorChan {
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(len(resultChan))
	for v := range resultChan {
		fmt.Println("Data received from service1:", v.GetMessage())
	}

	connSrv2, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer connSrv2.Close()

	serviceTwoResp, err := getDataFromServiceTwoClient(connSrv2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data received from service2:", serviceTwoResp.GetMessage())
}

func getDataFromServiceOneClient(ctx context.Context, conn *grpc.ClientConn) (*contract.Service1Response, error) {
	client1 := contract.NewGetDataFromService1Client(conn)
	log.Println("GetDataFromServiceOneClient")
	return client1.GetData(ctx, &contract.Service1Request{Id: 1})
}

func getDataFromServiceTwoClient(conn *grpc.ClientConn) (*contract.Service2Response, error) {
	client2 := contract.NewGetDataFromService2Client(conn)
	return client2.GetData(context.Background(), &contract.Service2Request{Id: 1})
}

func databaseOperation(repo repository.Repository) error {
	return repo.Insert(sampleData)
}
