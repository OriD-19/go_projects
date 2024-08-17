package database

import (
	"database/sql"
	"fmt"
	"log"

	"OriD19.com/auto_complete/data"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "contacts.db")

	if err != nil {
		panic("Error while opening the database")
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS contacts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name VARCHAR(255) NOT NULL,
            phone VARCHAR(30) NOT NULL
        );
        `)

	if err != nil {
		log.Fatal("Error while creating tables: " + err.Error())
	}

	DB = db
	populateAutoComplete()
}

// CreateContact takes a name and a phone and inserts a new row in the contacts table.
func CreateContact(name, phone string) error {
	_, err := DB.Exec("INSERT INTO contacts(name, phone) VALUES(?, ?)", name, phone)

	if err != nil {
		return err
	}

	// Update our auto-complete structure with the new name
	AutoCompleteTrie.Insert(name)

	return nil
}

// GetAllContacts retrieves all the rows from the database, and returns them into a slice.
func GetAllContacts() (*[]data.Contact, error) {
	rows, err := DB.Query("SELECT * FROM contacts;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []data.Contact
	for rows.Next() {
		var c data.Contact

		if err := rows.Scan(&c.Id, &c.Name, &c.Phone); err != nil {
			return nil, err
		}

		res = append(res, c)
	}

	return &res, nil
}

func GetContact(name string) (*data.Contact, error) {
	fmt.Println(name)
	row, err := DB.Query("SELECT * FROM contacts WHERE name = ?", name)

	if err != nil {
		return nil, err
	}

	var res data.Contact

	row.Next()
	if err := row.Scan(&res.Id, &res.Name, &res.Phone); err != nil {
		return nil, err
	}

	return &res, nil
}
