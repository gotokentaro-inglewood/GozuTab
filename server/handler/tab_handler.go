package handler

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gotokentaro-inglewood/GozuTab/repository"
)

func TabsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		tabs, err := repository.GetAllTabs(db)
		if err != nil {
			http.Error(w, "データ取得に失敗しました", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/tabs.html")
		if err != nil {
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, tabs)
	}
}

func CreateTabHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		userID, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			http.Error(w, "無効なuser_idです", http.StatusBadRequest)
			return
		}
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := repository.CreateTab(db, userID, title, content); err != nil {
			http.Error(w, "保存に失敗しました", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/tabs", http.StatusSeeOther)
	}
}

func UpdateTabHandler(db *sql.DB) http.HandlerFunc {
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
		title := r.FormValue("title")
		content := r.FormValue("content")

		if err := repository.UpdateTab(db, id, title, content); err != nil {
			http.Error(w, "更新に失敗しました", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteTabHandler(db *sql.DB) http.HandlerFunc {
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

		if err := repository.DeleteTab(db, id); err != nil {
			http.Error(w, "削除に失敗しました", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/tabs", http.StatusSeeOther)
	}
}
