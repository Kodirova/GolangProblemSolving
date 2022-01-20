package main

import (
	"KafkaTask/api/database"
	"KafkaTask/api/route"
	"KafkaTask/consumer"
	"fmt"

	"github.com/joho/godotenv"
)

var err error

func main() {
	godotenv.Load(".env")
	database.ConnectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println(database.DB == nil)
	go consumer.CreateContact()
	go consumer.UpdateContact()
	go consumer.DeleteContact()

	if err != nil {
		fmt.Println("Status:", err)
	}
	fmt.Println("we are here")
	route := route.SetUpRouter()
	route.Run(":9115")

}
