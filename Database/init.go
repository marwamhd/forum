package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	var err error
	DataBase.DB, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("error opening database:", err)
	}

	// Attempt to ping the database to verify connectivity
	err = DataBase.DB.Ping()
	if err != nil {
		log.Fatal("error pinging database:", err)
	}

	_, err = DataBase.DB.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	DataBase.CreateTable()

	// If no error occurred, the database connection is successfully established
	log.Println("Database connection established successfully")
}

func (DataBase *DB) CreateTable() {
	_, err := DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		session_id TEXT NULL
	)`)

	if err != nil {
		log.Fatal(err, "As")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		p_ID INTEGER NOT NULL,
		comment TEXT NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id)
	)`)

	if err != nil {
		log.Fatal(err, "s")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		title TEXT NOT NULL,
		post TEXT NOT NULL,
		cat1 INTEGER NOT NULL,
		cat2 INTEGER NOT NULL,
		cat3 INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id)

	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS author_liked_post (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		p_ID INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id)
	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
	_, err = DataBase.DB.Exec(`
	CREATE TABLE IF NOT EXISTS author_liked_comment (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		u_ID INTEGER NOT NULL,
		p_ID INTEGER NOT NULL,
		c_ID INTEGER NOT NULL,
		liked INTEGER NOT NULL,
		FOREIGN KEY (u_ID) REFERENCES users(id),
		FOREIGN KEY (p_ID) REFERENCES posts(id),
		FOREIGN KEY (c_ID) REFERENCES comments(id)
	)`)

	if err != nil {
		log.Fatal(err, "a")
	}
}

func (DataBase *DB) EmailQuery() error {
	rows, err := DataBase.DB.Query("select * from users")
	if err != nil {
		return err
	}
	var id int
	var email string
	var username string
	var password string

	for rows.Next() {
		err = rows.Scan(&id, &email, &username, &password)
		if err != nil {
			return err
		}
		fmt.Println(id)
		fmt.Println(email)
		fmt.Println(username)
		fmt.Println(password)
	}

	rows.Close()

	return nil
}

