package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Database of users
var Database *DB

func init() {

	db, err := sql.Open("sqlite3", "data/users.db")
	if err != nil {
		log.Println(err)
		log.Print("trying db/users.db")
		db, err = sql.Open("sqlite3", "db/users.db")
		if err != nil {
			log.Fatal(err)
		}
	}
	Database = &DB{db}
}

// DB implements CRUD for users db
type DB struct {
	*sql.DB
}

// User of db
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Get takes a first name and returns a user or error
func (db *DB) Get(firstName string) (*User, error) {
	stmt, err := db.Prepare("select * from users where first_name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	u := new(User)
	err = stmt.QueryRow(firstName).Scan(&u.ID, &u.FirstName, &u.LastName)
	return u, err
}

// AllUsers takes a first name and returns a user or error
func (db *DB) AllUsers() ([]*User, error) {
	stmt, err := db.Prepare("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	us := make([]*User, 0)
	for rows.Next() {
		u := new(User)
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName); err != nil {
			continue
		}
		us = append(us, u)
	}
	// f(&u.ID, &u.FirstName, &u.LastName)
	return us, nil
}
