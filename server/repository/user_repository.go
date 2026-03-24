package repository

import (
	"database/sql"

	"github.com/gotokentaro-inglewood/GozuTab/models"
)

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, email, icon_url FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.IconURL)
		users = append(users, u)
	}
	return users, nil
}

func UpdateUser(db *sql.DB, id int, name, iconURL string) error {
	_, err := db.Exec("UPDATE users SET name=$1, icon_url=$2 WHERE id=$3", name, iconURL, id)
	return err
}

func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
