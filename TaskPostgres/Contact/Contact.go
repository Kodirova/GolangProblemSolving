package Contact

import (
	"database/sql"
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

func CreateContact(db *sql.DB, c Contact) error {
	sqlStatement := `INSERT INTO contacts(firstname, lastname, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, c.FirstName, c.LastName, c.Phone, c.Email, c.Position).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateContact(db *sql.DB, c Contact, n int) int64 {
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

func DeleteContact(db *sql.DB, n int) error {
	_, err := db.Exec("DELETE from contacts where id=$1", n)
	fmt.Println(err)
	fmt.Println("Sucessfully deleted")
	return err
}

func GetContact(db *sql.DB, n int) (Contact, error) {
	var contact Contact
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(contact)
	return contact, err
}

func ListContacts(db *sql.DB) ([]Contact, error) {
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
