package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gotokentaro-inglewood/GozuTab/models"
	"github.com/gotokentaro-inglewood/GozuTab/repository"
)

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var user models.User
		err := db.QueryRow(
			`INSERT INTO users (name, email) VALUES ($1, $2) ON CONFLICT (email) DO NOTHING RETURNING id, name, email, COALESCE(icon_url, '')`,
			req.Name, req.Email,
		).Scan(&user.ID, &user.Name, &user.Email, &user.IconURL)
		if err != nil {
			http.Error(w, "保存に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "無効なIDです", http.StatusBadRequest)
			return
		}

		var req struct {
			Name    string `json:"name"`
			IconURL string `json:"icon_url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := repository.UpdateUser(db, id, req.Name, req.IconURL); err != nil {
			http.Error(w, "更新に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "updated"})
	}
}

func DeleteUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "無効なIDです", http.StatusBadRequest)
			return
		}

		if err := repository.DeleteUser(db, id); err != nil {
			http.Error(w, "削除に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
	}
}
