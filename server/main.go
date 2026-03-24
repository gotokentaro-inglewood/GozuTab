package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gotokentaro-inglewood/GozuTab/database"
	"github.com/gotokentaro-inglewood/GozuTab/handler"
	"github.com/gotokentaro-inglewood/GozuTab/repository"

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

	database.CreateUsersTable(db)
	database.CreateTabsTable(db)
	repository.InsertTestData(db)

	http.HandleFunc("/", handler.IndexHandler(db))
	http.HandleFunc("/users/create", handler.CreateUserHandler(db))
	http.HandleFunc("/users/update", handler.UpdateUserHandler(db))
	http.HandleFunc("/users/delete", handler.DeleteUserHandler(db))
	http.HandleFunc("/tabs", handler.TabsHandler(db))
	http.HandleFunc("/tabs/create", handler.CreateTabHandler(db))
	http.HandleFunc("/tabs/update", handler.UpdateTabHandler(db))
	http.HandleFunc("/tabs/delete", handler.DeleteTabHandler(db))

	fmt.Printf("Starting server on port %s...\n", appPort)
	addr := fmt.Sprintf("0.0.0.0:%s", appPort)
	fmt.Printf("Starting server at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}