package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gotokentaro-inglewood/GozuTab/models"
	"github.com/gotokentaro-inglewood/GozuTab/repository"
)

func TabsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tabs, err := repository.GetAllTabs(db)
		if err != nil {
			log.Printf("Error fetching tabs: %v\n", err)
			http.Error(w, "データ取得に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tabs)
	}
}

func CreateTabHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			UserID  int    `json:"user_id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		var tab models.Tab
		err := db.QueryRow(
			`INSERT INTO tabs (user_id, title, content) VALUES ($1, $2, $3) RETURNING id, user_id, title, COALESCE(artist,''), COALESCE(content,''), COALESCE(audio_url,''), COALESCE(status,'')`,
			req.UserID, req.Title, req.Content,
		).Scan(&tab.ID, &tab.UserID, &tab.Title, &tab.Artist, &tab.Content, &tab.AudioURL, &tab.Status)
		if err != nil {
			http.Error(w, "保存に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(tab)
	}
}

func UpdateTabHandler(db *sql.DB) http.HandlerFunc {
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
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if err := repository.UpdateTab(db, id, req.Title, req.Content); err != nil {
			http.Error(w, "更新に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "updated"})
	}
}

func DeleteTabHandler(db *sql.DB) http.HandlerFunc {
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

		if err := repository.DeleteTab(db, id); err != nil {
			http.Error(w, "削除に失敗しました", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
	}
}
