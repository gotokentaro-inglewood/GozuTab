package repository

import (
	"database/sql"

	"github.com/gotokentaro-inglewood/GozuTab/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}
