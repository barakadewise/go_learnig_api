package database;

import (
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)

func InitDd()*sql.DB{
	db,err :=sql.Open("sqlite3","./users.db");
	if err!=nil{
		log.Fatal(err);
	}
	
	//create users table if not exists
	log.Println("Creating users table if not exists...")
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		address TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL,
		dob TEXT NOT NULL,
		company_id INTERGER NOT NULL DEFAULT 1,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		
	);
	`
	_, err = db.Exec(createTableQuery)

	if err!=nil{
		log.Fatal(err);
	}
	log.Println("Users table created or already exists.")
	return db;



}