package main
import "testing"

func TestCreate(t *testing.T){
	contact := Contact{
		ID:        1,
		FirstName: "firstname",
		LastName:  "lastname",
		Phone:     "phone",
		Email:     "email",
		Position:  "postion",
	}
	got := createContact(contact)
	want :=  Contact{1, "firstname", "lastname", "phone", "email", "position"}
}