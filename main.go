package main

import (
	"adv/api"
	"log"
	"os"
)

func main() {
	logger()
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
