package Models

import (
	"crudTask/Database"
	"fmt"

	_ "github.com/lib/pq"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"

func CreateContact(contact *Contact) error {
	db := Database.ConnectDB()
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

func UpdateContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(tableContactCreation)
	if err1 != nil {
		return err1
	}
	sqlstatement := "UPDATE contacts SET firstname=$1, lastname=$2, phone=$3, email=$4, position=$5 WHERE id=$6"
	res, err := db.Exec(sqlstatement, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position, n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	fmt.Println(res)
	return nil
}

func DeleteContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(tableContactCreation)
	if err1 != nil {
		return err1
	}
	_, err := db.Exec("DELETE from contacts where id=$1", n)
	fmt.Println(res)
	fmt.Println("Sucessfully deleted")
	return err
}

func GetContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(tableContactCreation)
	if err1 != nil {
		return err1
	}
	sqlstatement := "SELECT id, firstname, lastname, phone, email, position FROM contacts WHERE id=$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(res)
	return err
}

func ListContacts(contact *[]Contact) error {
	db := Database.ConnectDB()
	res, err1 := db.Exec(tableContactCreation)
	if err1 != nil {
		return err1
	}
	rows, err := db.Query("SELECT * FROM contacts ORDER BY id")
	if err != nil {
		return err
	}
	fmt.Println(res)
	defer rows.Close()

	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Phone, &c.Email, &c.Position)
		fmt.Println(err)
		if err != nil {
			return err
		}
		*contact = append(*contact, c)

	}

	return nil
}
