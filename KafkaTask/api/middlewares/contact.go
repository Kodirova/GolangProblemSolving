package middlewares

import (
	"KafkaTask/api/database"
	"KafkaTask/api/model"

	"fmt"

	_ "github.com/lib/pq"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func CreateContact(contact *model.Contact) error {
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := database.DB.QueryRow(sqlStatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position).Scan(&contact.ID)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func UpdateContact(contact *model.Contact, n int) error {
	sqlstatement := "UPDATE contacts SET firstname=$1, lastname=$2, phone=$3, email=$4, position=$5 WHERE id=$6"
	fmt.Println(database.DB == nil)
	_, err := database.DB.Exec(sqlstatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position, n)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	return nil
}

func DeleteContact(contact *model.Contact, n int) error {
	_, err := database.DB.Exec("DELETE from contacts where id=$1", n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	fmt.Println("Sucessfully deleted")
	return err
}

func GetContact(contact *model.Contact, n int) error {
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := database.DB.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(contact)
	return err
}
