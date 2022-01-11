package main

import (
	"Structs/Contact"
	"fmt"

	_ "github.com/lib/pq"
)

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
		contact := Contact.Contact{
			FirstName: firstname,
			LastName:  lastname,
			Phone:     phone,
			Email:     email,
			Position:  postion,
		}
		fmt.Println(contact)
		Contact.contact.createContact()
		goto menu
	case 2:
		var search_id int
		fmt.Println("enter ID")
		fmt.Scanf("%d", &search_id)
		Contact.getContact(search_id)
		goto menu
	case 3:
		Contact.ListContacts()
		goto menu
	case 4:
		var delete_id int
		fmt.Println("enter ID")
		fmt.Scanf("%d", &delete_id)
		Contact.deleteContact(delete_id)
		goto menu
	case 5:
		var update_id int
		fmt.Println("enter ID")
		fmt.Scanf("%d", &update_id)
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
		contact := Contact.Contact{
			FirstName: firstname,
			LastName:  lastname,
			Phone:     phone,
			Email:     email,
			Position:  postion,
		}
		contact.updateContact(update_id)
		fmt.Println("after update")
		goto menu
	}
}
