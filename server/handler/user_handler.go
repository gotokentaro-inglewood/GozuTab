package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gotokentaro-inglewood/GozuTab/repository"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		query := `INSERT INTO users (name, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING;`
		_, err := db.Exec(query, name, email)
		if err != nil {
			http.Error(w, "保存に失敗しました", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "無効なIDです", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		iconURL := r.FormValue("icon_url")

		if err := repository.UpdateUser(db, id, name, iconURL); err != nil {
			http.Error(w, "更新に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "無効なIDです", http.StatusBadRequest)
			return
		}

		if err := repository.DeleteUser(db, id); err != nil {
			http.Error(w, "削除に失敗しました", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
