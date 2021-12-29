package structs

import "time"

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
