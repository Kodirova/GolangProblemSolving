package main

import (
	"crudTask/Database"
	"crudTask/Routes"
	"fmt"
)

var err error

func main() {
	Database.LoadEnv()
	Database.ConnectDB()

	if err != nil {
		fmt.Println("Status:", err)

	}
	route := Routes.SetUpRouter()
	route.Run(":8080")
}
