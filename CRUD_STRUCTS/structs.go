package main

import (
	"fmt"
	"time"
)

var ContactList map[int]Contact = make(map[int]Contact)
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

func createContact(c Contact) Contact {
	ContactList[c.ID] = c
	for key, value := range ContactList {
		fmt.Println("Key:", key, "Value:", value)
	}
	return c
}

func getAll() map[int]Contact {
	for key, value := range ContactList {
		fmt.Println("Key:", key, "Value:", value)
	}
	return ContactList
}

func get(n int) Contact {
	fmt.Println("we are here")
	value, exists := ContactList[n]
	fmt.Printf("key exists in map: %t, value: %v \n", exists, value)
	return ContactList[n]
}

func update(c Contact) Contact {
	ContactList[c.ID] = c
	c = Contact{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
		Email:     c.Email,
		Position:  c.Position,
	}
	return c
}

func deletebyId(n int) Contact {
	delete(ContactList, n)
	return ContactList[n]

}

func main() {
	var firstname, lastname, phone, email, postion string
menu:
	var choice int
	fmt.Println("MENU")
	fmt.Println("1.Create Contact")
	fmt.Println("2.Find Contact")
	fmt.Println("3. List of Contacts")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Update Contact")
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
	case 5:
		var search_id int
		fmt.Println("enter your id:")
		fmt.Scanf("%d", &search_id)
		if ContactList[search_id].FirstName != "" {
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
				ID:        search_id,
				FirstName: firstname,
				LastName:  lastname,
				Phone:     phone,
				Email:     email,
				Position:  postion,
			}
			fmt.Println("contact updated", contact)

		}
		goto menu
	}

}
