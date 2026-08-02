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
		log.Fatal("Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
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
	// Add in all the rows in the db here!!!
	rowsthatshouldbehere := "Email, Password"

	rows, err := db.Query(rowsthatshouldbehere)
	fmt.Printf(rowsthatshouldbehere)
	if err != nil {
		log.Printf("There are no Rows, replacing. \n")
	}
	if rows != nil {
		log.Fatalf("Emails exists? \n")
	}
}
