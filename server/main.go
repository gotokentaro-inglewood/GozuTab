package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベース接続の確認
	var dbString string
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			dbString = "success"
			fmt.Println("Database connection successful!")
			break
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Database connection status: %s", dbString)
	})

	fmt.Printf("Starting server on port %s...\n", appPort)
	addr := fmt.Sprintf("0.0.0.0:%s", appPort)
	fmt.Printf("Starting server at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}