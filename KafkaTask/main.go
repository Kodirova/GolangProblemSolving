package main

import (
	"KafkaTask/api/route"
	"KafkaTask/consumer"
	"fmt"
)

var err error

func main() {
	go consumer.CreateContact()
	if err != nil {
		fmt.Println("Status:", err)
	}
	fmt.Println("we are here")
	route := route.SetUpRouter()
	route.Run(":3030")

	fmt.Println("we are here")

}
