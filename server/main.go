package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"
	"github.com/gotokentaro-inglewood/GozuTab/models"
	"github.com/gotokentaro-inglewood/GozuTab/database"

	_ "github.com/lib/pq"
)

func main() {
	// 環境変数からデータベース接続情報を取得
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	appPort := os.Getenv("APP_PORT")
	dbPort := os.Getenv("DB_PORT")

	// データベース接続文字列を作成
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)

	// データベースに接続
	db := database.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
	defer db.Close()

	// テーブル作成のクエリ
	queryUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, -- ユーザーID（自動増分）
		name TEXT NOT NULL, -- ユーザー名
		email TEXT UNIQUE NOT NULL, -- メールアドレス（ユニーク制約）
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
		icon_url TEXT DEFAULT NULL -- アイコンのURLを保存するカラム
	);`

	_, err := db.Exec(queryUsers)
	if err != nil {
		log.Printf("Error creating users table: %v\n", err)
	} else {
		log.Println("Users table created successfully!")
	}

	queryTabs := `
	CREATE TABLE IF NOT EXISTS tabs (
		id SERIAL PRIMARY KEY, -- タブ譜ID（自動増分）
		user_id INTEGER REFERENCES users(id), -- 誰の投稿か（外部キー）
		title TEXT NOT NULL, -- タブ譜のタイトル
		artist TEXT, -- アーティスト名
		content TEXT,                          -- タブ譜のデータ
		audio_url TEXT,                         -- タブ譜に関連する音声ファイルのURL
		status TEXT,                              -- タブ譜の状態（例: "public", "private"）
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 作成日時
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- 更新日時
	);`

	_, err = db.Exec(queryTabs)
	if err != nil {
		log.Printf("Error creating tabs table: %v\n", err)
	} else {
		log.Println("Tabs table created successfully!")
	}

	// テストデータの挿入
	insertUser := `INSERT INTO users (name, email, icon_url) VALUES ($1, $2, $3) ON CONFLICT (email) DO NOTHING;`
	_, err = db.Exec(insertUser, "Kentaro", "kentaro@example.com", nil)
	if err != nil {
		log.Printf("Error inserting test user: %v\n", err)
	} else {
		log.Println("Test user inserted successfully!")
	}

	_, err = db.Exec(insertUser, "Yuki", "yuki@example.com", "https://example.com/yuki.jpg")
	if err != nil {
		log.Printf("Error inserting test user 2: %v\n", err)
	} else {
		log.Println("Test user 2 inserted successfully!")
	}

	insertTab := `INSERT INTO tabs (user_id, title, artist, content, audio_url, status) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err = db.Exec(insertTab, 1, "September", "Earth, Wind & Fire", "Tab content here", "https://example.com/september.mp3", "private")
	if err != nil {
		log.Printf("Error inserting test tab: %v\n", err)
	} else {
		log.Println("Test tab inserted successfully!")
	}

	_, err = db.Exec(insertTab, 1, "Magic Ways", "山下達郎", "Tab content here", "https://example.com/magic_ways.mp3", "public")
	if err != nil {
		log.Printf("Error inserting test tab 2: %v\n", err)
	} else {
		log.Println("Test tab 2 inserted successfully!")
	}

	_, err = db.Exec(insertTab, 2, "Plastic Love", "竹内まりや", "Tab content here", "https://example.com/plastic_love.mp3", "public")
	if err != nil {
		log.Printf("Error inserting test tab 3: %v\n", err)
	} else {
		log.Println("Test tab 3 inserted successfully!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// --- 1. ユーザー一覧 ---
		var users []models.User
		rows, _ := db.Query("SELECT name, email FROM users")
		defer rows.Close()
		for rows.Next() {
			var u models.User
			rows.Scan(&u.Name, &u.Email)
			users = append(users, u)
		}

		// --- 2. タブ譜一覧 ---
		var tabs []models.Tab
		tabRows, _ := db.Query("SELECT title FROM tabs")
		defer tabRows.Close()
		for tabRows.Next() {
			var t models.Tab
			tabRows.Scan(&t.Title)
			tabs = append(tabs, t)
		}

		// ページデータを準備
		data := models.PageData{
			Users: users,
			Tabs:  tabs,
		}

		// テンプレートを実行
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsing template: %v\n", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
	})

	fmt.Printf("Starting server on port %s...\n", appPort)
	addr := fmt.Sprintf("0.0.0.0:%s", appPort)
	fmt.Printf("Starting server at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}