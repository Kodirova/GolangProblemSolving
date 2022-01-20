package query

import (
	"kafkamicroservice/consumer/database"
	"kafkamicroservice/consumer/models"

	"fmt"

	_ "github.com/lib/pq"
)

func CreateContact(contact *models.Contact) error {
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := database.DB.QueryRow(sqlStatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position).Scan(&contact.ID)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func UpdateContact(contact *models.Contact, n int) error {
	sqlstatement := "UPDATE contacts SET firstname=$1, lastname=$2, phone=$3, email=$4, position=$5 WHERE id=$6"
	fmt.Println(database.DB == nil)
	_, err := database.DB.Exec(sqlstatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position, n)

	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}

	return nil
}

func DeleteContact(contact *models.Contact, n int) error {
	_, err := database.DB.Exec("DELETE from contacts where id=$1", n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	fmt.Println("Sucessfully deleted")
	return err
}

func GetContact(contact *models.Contact, n int) error {
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := database.DB.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(contact)
	return err
}

func ListContacts(contact *[]models.Contact) error {
	rows, err := database.DB.Query("SELECT * FROM contacts ORDER BY id")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var c models.Contact
		err = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Email, &c.Position)
		fmt.Println(err)
		if err != nil {
			return err
		}
		*contact = append(*contact, c)
	}
	return nil
}
