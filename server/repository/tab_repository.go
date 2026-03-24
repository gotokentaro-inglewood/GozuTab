package repository

import (
	"database/sql"

	"github.com/gotokentaro-inglewood/GozuTab/models"
)

func GetAllTabs(db *sql.DB) ([]models.Tab, error) {
	rows, err := db.Query("SELECT id, user_id, title, COALESCE(artist,''), COALESCE(content,''), COALESCE(audio_url,''), COALESCE(status,'') FROM tabs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tabs []models.Tab
	for rows.Next() {
		var t models.Tab
		rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Artist, &t.Content, &t.AudioURL, &t.Status)
		tabs = append(tabs, t)
	}
	return tabs, nil
}

func CreateTab(db *sql.DB, userID int, title, content string) error {
	_, err := db.Exec("INSERT INTO tabs (user_id, title, content) VALUES ($1, $2, $3)", userID, title, content)
	return err
}

func UpdateTab(db *sql.DB, id int, title, content string) error {
	_, err := db.Exec("UPDATE tabs SET title=$1, content=$2, updated_at=CURRENT_TIMESTAMP WHERE id=$3", title, content, id)
	return err
}

func DeleteTab(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM tabs WHERE id=$1", id)
	return err
}
