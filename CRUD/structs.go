package main

import (
	"fmt"
	"time"
)

var ContactList map[int]*Contact
var id int = 0

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

func createContact(c Contact) {
	ContactList[c.ID] = &c
	for key, value := range ContactList {
		fmt.Println("Key:", key, "Value:", value)
	}

}

func main() {
	ContactList = make(map[int]*Contact)
menu:
	var choice int
	fmt.Println("MENU")
	fmt.Println("Choose")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		var firstname, lastname, phone, email, postion string
		fmt.Println("enter first name")
		fmt.Scanf("%s", &firstname)
		fmt.Println("enter last name")
		fmt.Scanf("%s", &lastname)
		fmt.Println("enter phone")
		fmt.Scanf("%s", &phone)
		fmt.Println("enter email")
		fmt.Scanf("%s", &email)
		fmt.Println("enter postion")
		fmt.Scanf("%s", &postion)
		contact := Contact{
			ID:        id,
			FirstName: firstname,
			LastName:  lastname,
			Phone:     phone,
			Email:     email,
			Position:  postion,
		}
		id++
		fmt.Println("id:", id)
		fmt.Println("contact declared", contact)
		createContact(contact)
		goto menu
	}
}
