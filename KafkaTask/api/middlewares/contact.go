package middlewares

import (
	"KafkaTask/api/database"
	"KafkaTask/api/model"

	"fmt"

	_ "github.com/lib/pq"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func CreateContact(contact *model.Contact) error {
	db := database.ConnectDB()
	res, err1 := db.Exec(tableContactCreation)
	if err1 != nil {
		return err1
	}
	fmt.Println(res)
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position).Scan(&contact.ID)

	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil
}
