package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir? WILL BE USING ENV VARS NOW")
	}

	PORT := os.Getenv("BACKEND_API_PORT")
	dbstartup()
	startwebserver(PORT)

}
