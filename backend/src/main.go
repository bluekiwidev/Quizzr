package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	// Setup logger
	slog.SetDefault(logger)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir? WILL BE USING ENV VARS NOW")
	}

	// Get the port from environment variables
	PORT := os.Getenv("BACKEND_API_PORT")

	// Initialize the database connection
	dbstartup()
	defer dbAddr.Close()

	//Start API listener
	startwebserver(PORT)

}
