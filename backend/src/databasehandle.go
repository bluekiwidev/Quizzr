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

// This is not finished. Finish it later T_T
func tables(db *sql.DB) {
	// Add in all the Tables in the db here!!!
	query := "SHOW TABLES LIKE 'email'"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		fmt.Println("Table exists")
	} else {
		fmt.Println("Table does not exist")
		createtables(db)
	}
}

func createtables(db *sql.DB) {
	query := `CREATE TABLE email (id INT AUTO_INCREMENT PRIMARY KEY,email VARCHAR(255) NOT NULL)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Database failed to write: ", err)
	}
	fmt.Printf("\n Created Email Table")

	query = `CREATE TABLE password (id INT AUTO_INCREMENT PRIMARY KEY,password VARCHAR(255) NOT NULL)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Database failed to write: ", err)
	}
	fmt.Printf("\n Created Password Table")
}
