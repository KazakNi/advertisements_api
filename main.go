package main

import (
	"adv/api"
	database "adv/db"
	"log"
	"os"
)

var dev = os.Getenv("DEBUG") // false while production mode

func main() {
	logger()
	defer database.DB.Close()
	if dev == "true" {
		database.CreateDB(database.DB)
		database.ExecuteQueries(database.DB)
	}
	router := api.GetRoutes()
	router.Run(":8080")
}

func logger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
