package main

import (
	"kafkamicroservice/api/routes"
)

func main() {
	route := routes.SetUpRouter()
	route.Run(":8003")
}
