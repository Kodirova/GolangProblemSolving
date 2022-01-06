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

func getAll() {
	for key, value := range ContactList {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func get(n int) {
	fmt.Println("we are here")
	value, exists := ContactList[n]
	fmt.Printf("key exists in map: %t, value: %v \n", exists, value)
}

func main() {
	ContactList = make(map[int]*Contact)
	var firstname, lastname, phone, email, postion string
menu:
	var choice int
	fmt.Println("MENU")
	fmt.Println("1.Create Contact")
	fmt.Println("2.Find Contact")
	fmt.Println("3. List of Contacts")
	fmt.Println("4. Delete Contact")
	fmt.Println("Choose")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
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
	case 2:
		var search_id int
		fmt.Println("enter your id:")
		fmt.Scanf("%d", &search_id)
		get(search_id)
		goto menu
	case 3:
		getAll()
		goto menu
	case 4:
		var delete_id int
		fmt.Println("enter your id:")
		fmt.Scanf("%d", &delete_id)
		delete(ContactList, delete_id)
		fmt.Println("Succesfully deleted")
		goto menu
	}

}
