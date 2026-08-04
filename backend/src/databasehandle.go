package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func dbstartup() {
	fmt.Println("Starting DB")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPword := os.Getenv("DB_PWORD")
	dbIP := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")

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

// Column describes one column: its name, and the raw SQL that defines its
// type/constraints (everything after the name).
type Column struct {
	Name       string
	Definition string
}

// TableSchema describes one table and the columns it should have.
type TableSchema struct {
	Name    string
	Columns []Column
}

var expectedTables = []TableSchema{
	{
		Name: "users",
		Columns: []Column{
			{Name: "id", Definition: "INT AUTO_INCREMENT PRIMARY KEY"},
			{Name: "username", Definition: "VARCHAR(255) NOT NULL"},
		},
	},
	{
		Name: "userdata",
		Columns: []Column{
			{Name: "id", Definition: "INT AUTO_INCREMENT PRIMARY KEY"},
			{Name: "email", Definition: "VARCHAR(255) NOT NULL"},
			{Name: "username", Definition: "VARCHAR(255) NOT NULL"},
			{Name: "password", Definition: "VARCHAR(255) NOT NULL"},
			{Name: "datejoined", Definition: "DATETIME DEFAULT CURRENT_TIMESTAMP"},
			{Name: "isadmin", Definition: "BOOLEAN NOT NULL DEFAULT FALSE"},
		},
	},
}

func createtables(db *sql.DB, currentTables []string) {
	existing := make(map[string]bool, len(currentTables))
	for _, t := range currentTables {
		existing[t] = true
	}

	for _, table := range expectedTables {
		if existing[table.Name] {
			continue
		}

		defs := make([]string, len(table.Columns))
		for i, col := range table.Columns {
			defs[i] = col.Name + " " + col.Definition
		}

		query := fmt.Sprintf("CREATE TABLE %s (%s)", table.Name, strings.Join(defs, ", "))
		if _, err := db.Exec(query); err != nil {
			log.Fatalf("Database failed to write: %v", err)
		}
		fmt.Printf("\nCreated %s table\n", table.Name)
	}
}

func usernamevalidcheck(username string) int {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPword := os.Getenv("DB_PWORD")
	dbIP := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")

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
