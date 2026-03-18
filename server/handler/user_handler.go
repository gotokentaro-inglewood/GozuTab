package handler

import (
	"database/sql"
	"net/http"
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
