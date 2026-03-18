package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	
	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Println("Database connection successful!")
				fmt.Printf("Connected to database with DSN: %s\n", dsn)
				return db
			}
		}
		log.Printf("Retrying database connection... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Failed to connect to database")
	return nil
}

func CreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		icon_url TEXT DEFAULT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating users table: %v\n", err)
		return err
	}
	log.Println("Users table created successfully!")
	return nil
}

func CreateTabsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS tabs (
		id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES users(id),
		title TEXT NOT NULL,
		artist TEXT,
		content TEXT,
		audio_url TEXT,
		status TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error creating tabs table: %v\n", err)
		return err
	}
	log.Println("Tabs table created successfully!")
	return nil
}