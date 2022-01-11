package Models

import (
	"crudTask/Database"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateContact(contact *Contact) error {
	db := Database.ConnectDB()
	sqlStatement := `INSERT INTO contact(first_name, last_name, phone, email, position) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sqlStatement, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position).Scan(&contact.ID)

	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func UpdateContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	sqlstatement := "UPDATE contact SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6"
	res, err := db.Exec(sqlstatement, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position, n)
	if err != nil {
		fmt.Printf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return nil
}

func DeleteContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	_, err := db.Exec("DELETE from contact where id=$1", n)
	fmt.Println(err)
	fmt.Println("Sucessfully deleted")
	return err
}

func GetContact(contact *Contact, n string) error {
	db := Database.ConnectDB()
	sqlstatement := "SELECT id, first_name, last_name, phone, email, position FROM contact WHERE id=$1"
	row := db.QueryRow(sqlstatement, n)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	fmt.Println(err)
	return err
}

func ListContacts(contact *[]Contact) error {
	db := Database.ConnectDB()

	rows, err := db.Query("SELECT * FROM contact ORDER BY id")
	if err != nil {
		return err
	}

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
