package Contact

import (
	"Structs/Database"
	"fmt"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func (c Contact) createContact() error {
	db := Database.LoadDb()
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, c.FirstName, c.LastName, c.Phone, c.Email, c.Position).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c Contact) updateContact(n int) int64 {
	db := Database.LoadDb()
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
	db := Database.LoadDb()
	_, err := db.Exec("DELETE from contacts where id=$1", n)
	fmt.Println(err)
	fmt.Println("Sucessfully deleted")
	return err
}

func getContact(n int) (Contact, error) {
	db := Database.LoadDb()
	var contact Contact
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(contact)
	return contact, err
}

func ListContacts() ([]Contact, error) {
	db := Database.LoadDb()
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
