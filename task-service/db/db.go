package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
		return
	}

	cnStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"))

	DB, err = sql.Open("postgres", cnStr)

	if err != nil {
		log.Fatal("Failed to connect to postgres: ", err)
		return
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error Pinging database!")
		return
	}

	fmt.Println("Task Service Connected to Database Hehehe!")

}
