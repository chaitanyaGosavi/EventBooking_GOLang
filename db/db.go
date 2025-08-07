package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDBConnection() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Database connection failed!!!!")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	fmt.Println("DataBase Connectin Successful")

	CreateEventsTable()
}

func CreateEventsTable() {
	var createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTableQuery)
	if err != nil {
		panic("Users Table Creation failed !!!!!!!")
	}

	var createEventsTableQuery = `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    datetime DATETIME NOT NULL,
    userid INTEGER NOT NULL,
	FOREIGN KEY(userid) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTableQuery)

	if err != nil {
		panic("Events Table Creation failed !!!!!!!")
	}

	var createRegistrationsTableQuery = `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER
	user_id INTEGER
	FOREIGN KEY(event_id) REFERENCES events(id)
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTableQuery)

	if err != nil {
		panic("Events Table Creation failed !!!!!!!")
	}
}
