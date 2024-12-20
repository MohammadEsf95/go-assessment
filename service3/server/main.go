package main

import (
	"log"
	"service3/infrastructure/database"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// implement GetResult

	// go GetData from service1

	// query operation

	// if err != nil in two ops, then dont call GetData from service 2

	//implement test
}
