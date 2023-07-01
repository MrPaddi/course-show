package models

// import sqlite driver
import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Database is a struct to store database connection
type Database struct {
	*sql.DB
}

// InitDB is a function to initialize database connection
func InitDB(filepath string) *Database {
	// check if file exist
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		log.Fatal("Database not found")
	}

	// open database
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	// create table if not exist
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS account (id INTEGER PRIMARY KEY, username TEXT, email TEXT, password TEXT, phone TEXT, is_active BOOLEAN, is_delete BOOLEAN, created INTEGER, updated INTEGER, first_login BOOLEAN)")
	if err != nil {
		log.Fatal(err)
	}

	// return database connection
	return &Database{db}
}

// Account is a struct to store account data for sqlite3 database, need ID, Username, Email, Password, Phone, IsActive, IsDelete, Created, Updated, FirstLogin
// type Account struct {
// 	ID         int    `json:"id"`
// 	Username   string `json:"username"`
// 	Email      string `json:"email"`
// 	Password   string `json:"password"`
// 	Phone      string `json:"phone"`
// 	IsActive   bool   `json:"is_active"`
// 	IsDelete   bool   `json:"is_delete"`
// 	Created    int64  `json:"created"`
// 	Updated    int64  `json:"updated"`
// 	FirstLogin bool   `json:"first_login"`
// }
