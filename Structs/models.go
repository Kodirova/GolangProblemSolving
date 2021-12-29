package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  int
	CreatedAt time.Time
	CreatedBy time.Time
	DueDate   time.Time
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "iroda1808"
	dbname   = "test"
)

func main() {
	//connection to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
}
