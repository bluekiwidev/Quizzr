package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
	}

	PORT := os.Getenv("WEBSERVER_PORT")

	dbstartup()
	startwebserver(PORT)
}
