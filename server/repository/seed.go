package repository

import (
	"database/sql"
	"log"
)

func InsertTestData(db *sql.DB) {
	insertUser := `INSERT INTO users (name, email, icon_url) VALUES ($1, $2, $3) ON CONFLICT (email) DO NOTHING;`

	_, err := db.Exec(insertUser, "Kentaro", "kentaro@example.com", nil)
	if err != nil {
		log.Printf("Error inserting test user: %v\n", err)
	} else {
		log.Println("Test user inserted successfully!")
	}

	_, err = db.Exec(insertUser, "Yuki", "yuki@example.com", "https://example.com/yuki.jpg")
	if err != nil {
		log.Printf("Error inserting test user 2: %v\n", err)
	} else {
		log.Println("Test user 2 inserted successfully!")
	}

	insertTab := `INSERT INTO tabs (user_id, title, artist, content, audio_url, status) VALUES ($1, $2, $3, $4, $5, $6);`

	_, err = db.Exec(insertTab, 1, "September", "Earth, Wind & Fire", "Tab content here", "https://example.com/september.mp3", "private")
	if err != nil {
		log.Printf("Error inserting test tab: %v\n", err)
	} else {
		log.Println("Test tab inserted successfully!")
	}

	_, err = db.Exec(insertTab, 1, "Magic Ways", "山下達郎", "Tab content here", "https://example.com/magic_ways.mp3", "public")
	if err != nil {
		log.Printf("Error inserting test tab 2: %v\n", err)
	} else {
		log.Println("Test tab 2 inserted successfully!")
	}

	_, err = db.Exec(insertTab, 2, "Plastic Love", "竹内まりや", "Tab content here", "https://example.com/plastic_love.mp3", "public")
	if err != nil {
		log.Printf("Error inserting test tab 3: %v\n", err)
	} else {
		log.Println("Test tab 3 inserted successfully!")
	}
}
