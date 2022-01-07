package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type Task struct {
	ID        int
	Name      string
	Status    string
	Priority  int
	CreatedAt time.Time
	CreatedBy time.Time
	DueDate   time.Time
}

const (
	host     = "localhost"
	port     = 5432
	user     = "Kimmy"
	password = "Kimmy@1808"
	dbname   = "test"
)

const tableContactCreation = "CREATE TABLE IF NOT EXISTS contacts(id SERIAL, firstname TEXT NOT NULL, lastname TEXT NOT NULL, phone VARCHAR(13), email text, position text)"
const tableTaskCreation = "CREATE TABLE IF NOT EXISTS tasks"

func (c *Contact) createContact(db *sql.DB) error {
	err := db.QueryRow("INSERET INTO contacts (id, firstname, lastname, phone, email, position) VALUES ($1, $2, $3, $4, $5, $6) returning id", c.FirstName, c.LastName, c.Phone, c.Email, c.Position).Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Contact) updateContact(db *sql.DB) error {
	_, err := db.Exec("UPDATE contacts SET firstname=$1, lastname=$2, phone=$3, email=$4, position=$5 WHERE id=$6",
		c.FirstName, c.LastName, c.Phone, c.Email, c.Position, c.ID)
	return err
}

func (c *Contact) deleteContact(db *sql.DB) error {
	_, err := db.Exec("DELETE from contacts where id=$1", c.ID)
	return err
}

func (c *Contact) getContact(db *sql.DB) error {
	return db.QueryRow("SELECT id, firstname, lastname, phone, email, position from contacts where id=$1", c.ID).Scan(&c.FirstName, c.LastName, c.Phone, c.Email, c.Position)
}

func (c *Contact) ListContacts(db *sql.DB) ([]Contact, error) {
	contacts := []Contact{}
	rows, err := db.Query("SELECT id, firstname, lastname, phone, email, position FROM contacts ORDER BY id")
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
	return contacts, nil
}

func LoadDb() *sql.DB {
	//connection to database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")
	return db
}
func main() {

}
