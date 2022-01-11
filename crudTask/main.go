package main

import (
	"crudTask/Database"
	"crudTask/Routes"
	"fmt"
)

var err error

func main() {
	Database.ConnectDB()
	if err != nil {
		fmt.Println("Status:", err)

	}
	route := Routes.SetUpRouter()
	route.Run()
}
