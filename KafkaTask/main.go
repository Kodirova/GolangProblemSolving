package main

import (
	"KafkaTask/api/database"
	"KafkaTask/api/route"
	"KafkaTask/consumer"
	"fmt"
)

var err error

func main() {
	route := route.SetUpRouter()
	route.Run(":8080")
	database.LoadEnv()
	database.ConnectDB()

	if err != nil {
		fmt.Println("Status:", err)

	}
	consumer.CreateContact()
	if err != nil {
		fmt.Println("Status:", err)

	}

}
