package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "./db/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Succesfully connected to Sqlite")
}

func CreateDB() {
	DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		first_name VARCHAR(100),
		last_name VARCHAR(100),
		display_email VARCHAR(100),
		password VARCHAR(100) NOT NULL,
		pic_path VARCHAR(100),
		link VARCHAR(100)
	)`)

	DB.Exec(`CREATE TABLE IF NOT EXISTS links (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		url VARCHAR(100) NOT NULL,
		platform VARCHAR(50) NOT NULL,
		order_index INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id)
	);`)
}
