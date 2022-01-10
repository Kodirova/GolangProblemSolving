package main

import (
	"crudTask/Database"
	"crudTask/Models"
	"crudTask/Routes"
	"fmt"
)

var err error

func main() {
	Database.ConnectDB()
	Database.DB.AutoMigrate(&Models.Contact{})
	if err != nil {
		fmt.Println("Status:", err)

	}
	route := Routes.SetUpRouter()
	route.Run()
}
