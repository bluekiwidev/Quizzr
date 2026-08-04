package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var dbName = os.Getenv("DB_NAME")
var dbUser = os.Getenv("DB_USER")
var dbPword = os.Getenv("DB_PWORD")
var dbIP = os.Getenv("DB_IP")
var dbPort = os.Getenv("DB_PORT")

func dbstartup() {
	fmt.Println("Starting DB")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir? WILL BE USING ENV VARS")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?parseTime=true", dbUser, dbPword, dbIP, dbPort, dbName)
	fmt.Println(dsn, "\n")

	db, err := sql.Open("mysql", dsn) // Initializing
	if err != nil {
		log.Fatalf("\n Failed to Make a Opening with database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("\n Failed to Make a Connection to database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}

	fmt.Println("Database Connected Successfully!")
	tables(db)
}

// This is not finished. Finish it later T_T
func tables(db *sql.DB) {
	query := "SHOW TABLES"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer rows.Close()

	var currentTables []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalf(err.Error())
		}
		currentTables = append(currentTables, name)
	}
	createtables(db, currentTables)
}

func createtables(db *sql.DB, currentTables []string) {
	// ADD ALL THE TABLES THAT NEEDS TO BE HERE BELOW.
	Tablesthatshouldbehere := []string{"usernames", "passwords"}

	for _, required := range Tablesthatshouldbehere {
		found := false
		for _, existing := range currentTables {
			if required == existing {
				found = true
				break
			}
		}

		if !found {
			query := fmt.Sprintf("CREATE TABLE %s (id INT AUTO_INCREMENT PRIMARY KEY, %s VARCHAR(255) NOT NULL)", required, required)
			_, err := db.Exec(query)
			if err != nil {
				log.Fatalf("Database failed to write: %v", err)
			}
			fmt.Printf("\nCreated %s table\n", required)
		}
	}
}

func usernamevalidcheck(username string) int {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir? WILL BE USING ENV VARS")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?parseTime=true", dbUser, dbPword, dbIP, dbPort, dbName)
	fmt.Println(dsn, "\n")

	db, err := sql.Open("mysql", dsn) // Initializing
	if err != nil {
		log.Fatalf("\n Failed to Make a Opening with database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("\n Failed to Make a Connection to database. HINT: Are your .env credentials correct?                   ERROR:", err)
	}

	fmt.Println("Database Connected Successfully!")

	var exists bool

	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM usernames WHERE usernames = ?)", username).Scan(&exists)

	if err != nil {
		fmt.Println("QueryRow error:", err)
		return 500
	}

	if exists {
		fmt.Println("Sent code 409")
		return 409 // username taken
	} else {
		fmt.Println("Sent code 200")
		return 200 // username available
	}
}
