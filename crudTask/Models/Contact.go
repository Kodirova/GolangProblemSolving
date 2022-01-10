package Models

import (
	"crudTask/Database"
	_ "fmt"

	_ "github.com/lib/pq"
)

func CreateContact(contact *Contact) (err error) {
	if err = Database.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}
