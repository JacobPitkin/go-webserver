package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"jacobpitkin.com/webserv/database"
)

const DEFAULT_LOCATION = "./default.db"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbLocation := os.Getenv("DB")

	if dbLocation == "" {
		dbLocation = DEFAULT_LOCATION
	}

	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("opened db")

	database.CreateDb(db)

	database.InsertUser(db)

	// database.SelectUser(db)

	database.SelectUsers(db)

	// database.DeleteUser(db)

	// database.DeleteAllUsers(db)
}
