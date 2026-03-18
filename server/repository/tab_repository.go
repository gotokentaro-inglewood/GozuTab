package repository

import (
	"database/sql"

	"github.com/gotokentaro-inglewood/GozuTab/models"
)

func GetAllTabs(db *sql.DB) ([]models.Tab, error) {
	rows, err := db.Query("SELECT title FROM tabs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tabs []models.Tab
	for rows.Next() {
		var t models.Tab
		rows.Scan(&t.Title)
		tabs = append(tabs, t)
	}
	return tabs, nil
}
