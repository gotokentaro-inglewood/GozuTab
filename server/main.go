package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gotokentaro-inglewood/GozuTab/database"
	"github.com/gotokentaro-inglewood/GozuTab/handler"
	"github.com/gotokentaro-inglewood/GozuTab/repository"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

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

	auth := handler.AuthMiddleware
	cors := handler.CORSMiddleware
	http.HandleFunc("/api/health", cors(handler.HealthHandler()))
	http.HandleFunc("/", cors(auth(handler.IndexHandler(db))))
	http.HandleFunc("/users/create", cors(auth(handler.CreateUserHandler(db))))
	http.HandleFunc("/users/update", cors(auth(handler.UpdateUserHandler(db))))
	http.HandleFunc("/users/delete", cors(auth(handler.DeleteUserHandler(db))))
	http.HandleFunc("/tabs", cors(auth(handler.TabsHandler(db))))
	http.HandleFunc("/tabs/create", cors(auth(handler.CreateTabHandler(db))))
	http.HandleFunc("/tabs/update", cors(auth(handler.UpdateTabHandler(db))))
	http.HandleFunc("/tabs/delete", cors(auth(handler.DeleteTabHandler(db))))

	fmt.Printf("Starting server on port %s...\n", appPort)
	addr := fmt.Sprintf("0.0.0.0:%s", appPort)
	fmt.Printf("Starting server at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}