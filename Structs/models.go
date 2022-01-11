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
	user     = "Kimmy"
	password = "Kimmy@1808"
	dbname   = "test"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func (c Contact) createContact() error {
	db := LoadDb()
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, c.FirstName, c.LastName, c.Phone, c.Email, c.Position).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c Contact) updateContact(n int) int64 {
	db := LoadDb()
	fmt.Println("We are in update")
	sqlstatement := "UPDATE contacts SET firstname=$1, lastname=$2, phone=$3, email=$4, position=$5 WHERE id=$6"
	res, err := db.Exec(sqlstatement, c.FirstName, c.LastName, c.Phone, c.Email, c.Position, n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteContact(n int) error {
	db := LoadDb()
	_, err := db.Exec("DELETE from contacts where id=$1", n)
	fmt.Println(err)
	fmt.Println("Sucessfully deleted")
	return err
}

func getContact(n int) (Contact, error) {
	db := LoadDb()
	var contact Contact
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(contact)
	return contact, err
}

func ListContacts() ([]Contact, error) {
	db := LoadDb()
	var contacts []Contact
	rows, err := db.Query("SELECT * FROM contacts ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Email, &c.Position)
		if err != nil {
			return contacts, err
		}
		contacts = append(contacts, c)
	}
	fmt.Println(contacts)
	return contacts, nil
}

func LoadDb() *sql.DB {
	//connection to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
	return db
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
			FirstName: firstname,
			LastName:  lastname,
			Phone:     phone,
			Email:     email,
			Position:  postion,
		}
		fmt.Println(contact)
		contact.createContact()
		goto menu
	case 2:
		var search_id int
		fmt.Println("enter ID")
		fmt.Scanf("%d", &search_id)
		getContact(search_id)
		goto menu
	case 3:
		ListContacts()
		goto menu
	case 4:
		var delete_id int
		fmt.Println("enter ID")
		fmt.Scanf("%d", &delete_id)
		deleteContact(delete_id)
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
		contact := Contact{
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
