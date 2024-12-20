package main

import (
	"log"
	"service3/entity"
	"service3/infrastructure/database"
	"service3/repository"
	"time"
)

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

	err = db.AutoMigrate(&entity.Model{})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	err = repo.Insert(sampleData)
	if err != nil {
		log.Fatal(err)
	}

	// implement GetResult

	// go GetData from service1

	// query operation

	// if err != nil in two ops, then dont call GetData from service 2

	//implement test
}
