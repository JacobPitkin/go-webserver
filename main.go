package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
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

	createDb(db)

	insertUser(db)

	// selectUser(db)

	selectUsers(db)

	// deleteUser(db)

	// deleteAllUsers(db)
}

func createDb(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("created table (if it didn't already exist anyways)")
}

func insertUser(db *sql.DB) {
	username := "Jacob"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users(username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	fmt.Printf("inserted user with id: %d", id)
}

func selectUser(db *sql.DB, userId int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, userId).Scan(&id, &username, &password, &createdAt); err != nil {
		log.Fatal(err)
	}

	fmt.Println(id, username, password, createdAt)
}

func selectUsers(db *sql.DB) {
	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user

		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", users)
}

func deleteUser(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("deleted user id %d\n", id)
}

func deleteAllUsers(db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted all users")
}
