package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gotokentaro-inglewood/GozuTab/models"
	"github.com/gotokentaro-inglewood/GozuTab/repository"
)

func IndexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			log.Printf("Error fetching users: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		tabs, err := repository.GetAllTabs(db)
		if err != nil {
			log.Printf("Error fetching tabs: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := models.PageData{
			Users: users,
			Tabs:  tabs,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
