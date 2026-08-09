package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var dbAddr *sql.DB

func dbstartup() {
	//Stat MariaDB conn
	logger.Info("Starting MariaDB connection")
	err := godotenv.Load()
	if err != nil {
		logger.Warn("Could not find .env file. Will use env vars instead")
	}

	//Declare DB connection variables
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPword := os.Getenv("DB_PWORD")
	dbIP := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?parseTime=true", dbUser, dbPword, dbIP, dbPort, dbName)
	logger.Info(fmt.Sprint("MariaDB DSN is: ", dsn))

	// Connect to the database
	db, err := sql.Open("mysql", dsn) // Initializing
	if err != nil {
		logger.Error("Failed to open database connection", "error", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	//Set global dbAddr variable to the database connection+
	dbAddr = db

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprint("Failed to ping MariaDB, error: ", err))
	}

	logger.Info("MariaDB connection successful")
	tables(db)
}

func redisinit() (*redis.Client, context.Context) {
	logger.Info("Starting Redis")

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Could not find .env file. Will use env vars instead")
	}

	// Setup redis parameters
	redis_addr := os.Getenv("REDIS_ADDR")
	redis_pword := os.Getenv("REDIS_PWORD")
	redis_db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		logger.Error(fmt.Sprintf("Env var REDIS_DB is required and must be a valid number: %v", err))
	}
	redis_protocol, err := strconv.Atoi(os.Getenv("REDIS_PROTOCOL"))
	if err != nil {
		logger.Error(fmt.Sprintf("Env var REDIS_PROTOCOL is required and must be a valid number: %v", err))
	}

	// Connect to database
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pword,
		DB:       redis_db,
		Protocol: redis_protocol,
	})

	ctx := context.Background()

	// Test connection
	err = rdb.Ping(ctx).Err()
	if err != nil {
		logger.Error("Failed to ping Redis")
	}

	logger.Info("Redis connection successful")
	return rdb, ctx
}

// This is not finished. Finish it later T_T
func tables(db *sql.DB) {
	query := "SHOW TABLES"

	rows, err := db.Query(query)
	if err != nil {
		logger.Error(err.Error())
	}
	defer rows.Close()

	var currentTables []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			logger.Error(err.Error())
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
			logger.Error(fmt.Sprintf("Database failed to write: %v", err))
		}
		logger.Info(fmt.Sprintf("\nCreated %s table\n", table.Name))
	}
}

func usernamevalidcheck(username string) int {
	logger.Info(fmt.Sprint("Checking username: ", username))
	var exists bool

	err := dbAddr.QueryRow("SELECT EXISTS(SELECT 1 FROM userdata WHERE username = ?)", username).Scan(&exists)

	if err != nil {
		logger.Error(fmt.Sprint("QueryRow error:", err))
		return 500
	}

	if exists {
		logger.Info("Returned 409, taken")
		return 409 // username taken
	} else {
		logger.Info("Returned 200, avalible")
		return 200 // username available
	}
}

func adduser(username string, email string, password string) bool {
	query := "INSERT INTO userdata (username, email, password) VALUES (?, ?, ?)"

	_, err := dbAddr.Exec(query, username, email, password)
	if err != nil {
		logger.Error(fmt.Sprint("Error adding user to userdata, error: ", err))
		return false
	}

	return true
}

func compareemail(email string) bool {
	var exists bool
	err := dbAddr.QueryRow("SELECT EXISTS(SELECT 1 FROM userdata WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		logger.Error(fmt.Sprint("Error checking email, error: ", err))
		return false
	}
	return (exists)

}

func comparepassword(email string, password string) bool {
	var storedPassword string
	err := dbAddr.QueryRow("SELECT password FROM userdata WHERE email = ?", email).Scan(&storedPassword)
	if err != nil {
		return false
	}

	if bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)) == nil {
		return true
	}
	return false
}
