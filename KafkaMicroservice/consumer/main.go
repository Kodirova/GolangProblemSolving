package main

import (
	"kafkamicroservice/consumer/database"

	"kafkamicroservice/consumer/worker"

	"kafkamicroservice/consumer/route"

	"github.com/joho/godotenv"
)

var err error

func main() {
	godotenv.Load(".env")
	database.ConnectDB()
	if err != nil {
		panic(err)
	}
	go worker.CreateContact()
	go worker.UpdateContact()
	go worker.DeleteContact()
	route := route.SetUpRouter()
	route.Run(":8085")

}
