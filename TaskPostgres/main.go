package main

import (
	"TaskPostgres/Config"
	"TaskPostgres/Contact"
	"TaskPostgres/Task"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	DB = Config.ConnectDB()
	var firstname, lastname, phone, email, postion string
	var name, status, priority, createdby string

main_menu:
	var main_choice int
	fmt.Println("Choose")
	fmt.Println("1 menu for contact")
	fmt.Println("2 menu for task")
	fmt.Scanf("%d", &main_choice)
	switch main_choice {
	case 1:
		fmt.Println("menu for contact")
	menu1:
		var choice int
		fmt.Println("MENU")
		fmt.Println("1.Create Contact")
		fmt.Println("2.Find Contact")
		fmt.Println("3. List of Contacts")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Update Contact")
		fmt.Println("6. Main Menu")
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
			Contact.CreateContact(DB, contact)
			goto menu1
		case 2:
			var search_id int
			fmt.Println("enter ID")
			fmt.Scanf("%d", &search_id)
			Contact.GetContact(DB, search_id)
			goto menu1
		case 3:
			Contact.ListContacts(DB)
			goto menu1
		case 4:
			var delete_id int
			fmt.Println("enter ID")
			fmt.Scanf("%d", &delete_id)
			Contact.DeleteContact(DB, delete_id)
			goto menu1
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
			Contact.UpdateContact(DB, contact, update_id)
			fmt.Println("after update")
			goto menu1
		case 6:
			goto main_menu
		}
	case 2:
	menu2:
		var choice int
		fmt.Println("MENU")
		fmt.Println("1.Create Task")
		fmt.Println("2.Find Task")
		fmt.Println("3. List of Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Update Task")
		fmt.Println("6. Main Menu")
		fmt.Println("Choose")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			fmt.Println("enter  name")
			fmt.Scanf("%s", &name)
			fmt.Println("enter status")
			fmt.Scanf("%s", &status)
			fmt.Println("enter priority")
			fmt.Scanf("%s", &priority)
			fmt.Println("enter createdby")
			fmt.Scanf("%s", &createdby)

			task := Task.Task{
				Name:      name,
				Status:    status,
				Priority:  priority,
				CreatedBy: createdby,
			}
			fmt.Println(task)
			Task.CreateTask(DB, task)
			fmt.Println("2")
			goto menu2
		case 2:
			var search_id int
			fmt.Println("enter ID")
			fmt.Scanf("%d", &search_id)
			Task.GetTask(DB, search_id)
			goto menu2
		case 3:
			Task.ListTasks(DB)
			goto menu2
		case 4:
			var delete_id int
			fmt.Println("enter ID")
			fmt.Scanf("%d", &delete_id)
			Task.DeleteTask(DB, delete_id)
			goto menu2
		case 5:
			var update_id int
			fmt.Println("enter ID")
			fmt.Scanf("%d", &update_id)
			fmt.Println("enter  name")
			fmt.Scanf("%s", &name)
			fmt.Println("enter status")
			fmt.Scanf("%s", &status)
			fmt.Println("enter priority")
			fmt.Scanf("%s", &priority)
			fmt.Println("enter createdby")
			fmt.Scanf("%s", &createdby)

			task := Task.Task{
				Name:      name,
				Status:    status,
				Priority:  priority,
				CreatedBy: createdby,
			}
			fmt.Println(task)
			Task.UpdateTask(DB, task, update_id)
			fmt.Println("3")
			goto menu2
		case 6:
			goto main_menu
		}
	default:
		break
	}
}
