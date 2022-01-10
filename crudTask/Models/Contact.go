package Models

import (
	"crudTask/Database"
	"fmt"
	_ "fmt"

	_ "github.com/lib/pq"
)

func CreateContact(contact *Contact) (err error) {
	if err = Database.DB.Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func ListContacts(contact *[]Contact) (err error) {
	if err = Database.DB.Find(contact).Error; err != nil {
		return err
	}
	return nil
}

func GetContact(contact *Contact, id string) (err error) {
	if err = Database.DB.Where("id = ?", id).First(contact).Error; err != nil {
		return err
	}
	return nil
}

func UpdateContact(contact *Contact, id string) (err error) {
	fmt.Println(contact)
	Database.DB.Save(contact)
	return nil
}

func DeleteContact(contact *Contact, id string) (err error) {
	Database.DB.Where("id = ?", id).Delete(contact)
	return nil
}
