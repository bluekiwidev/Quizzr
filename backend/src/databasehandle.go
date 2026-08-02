package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func dbstartup() {
	fmt.Println("Starting DB")
	// .env read
	err := godotenv.Load()
	if err != nil {
		log.Println("Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir? Will try to use system env variables instead...", err)
	}

	DB := os.Getenv("DB")

	// The db actually starts startinging
	dsn := DB // The maria db connection

	db, err := sql.Open("mysql", dsn) // Initializing

	if err != nil {
		log.Fatalf("Failed to Make a Opening with database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to Make a Connection to database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}

	fmt.Println("Database Connected Successfully!")
	tables(db)
}

func tables(db *sql.DB) {
	// Add in all the Tables in the db here!!!
	tablesthatshouldbehere := "email, password"

	rows, err := db.Query(tablesthatshouldbehere)
	fmt.Printf(tablesthatshouldbehere)
	if err != nil {
		log.Printf("\n There are no Table, replacing.")
	}
	if rows != nil {
		log.Fatalf("\n Emails exists?")
	}
}
