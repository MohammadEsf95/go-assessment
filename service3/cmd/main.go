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

	connSrv2, err := grpc.NewClient(":8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer connSrv2.Close()

	serviceOneClient := contract.NewGetDataFromService1Client(connSrv1)
	serviceTwoClient := contract.NewGetDataFromService2Client(connSrv2)

	GetResult(repo, serviceOneClient, serviceTwoClient)
}

func GetResult(repo repository.Repository,
	srvOneClient contract.GetDataFromService1Client,
	srvTwoClient contract.GetDataFromService2Client) {
	var wg sync.WaitGroup
	errorChan := make(chan error, 2)
	resultChan := make(chan *contract.Service1Response, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		serviceOneResp, err := getDataFromServiceOneClient(context.Background(), srvOneClient)
		if err != nil {
			errorChan <- fmt.Errorf("getDataFromServiceOneClient error: %w", err)
			return
		}
		resultChan <- serviceOneResp
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := databaseOperation(repo); err != nil {
			errorChan <- fmt.Errorf("databaseOperation error: %w", err)
		}
		log.Println("databaseOperation done")
	}()

	wg.Wait()
	close(resultChan)
	close(errorChan)

	for err := range errorChan {
		if err != nil {
			log.Fatal(err)
		}
	}

	for v := range resultChan {
		fmt.Println("Data received from service1:", v.GetMessage())
	}

	resultFromServiceTwo, err := getDataFromServiceTwoClient(srvTwoClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data received from service2:", resultFromServiceTwo.GetMessage())
}

func getDataFromServiceOneClient(ctx context.Context, client contract.GetDataFromService1Client) (*contract.Service1Response, error) {
	return client.GetData(ctx, &contract.Service1Request{Id: 1})
}

func getDataFromServiceTwoClient(client contract.GetDataFromService2Client) (*contract.Service2Response, error) {
	return client.GetData(context.Background(), &contract.Service2Request{Id: 1})
}

func databaseOperation(repo repository.Repository) error {
	return repo.Insert(sampleData)
}
