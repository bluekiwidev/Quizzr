package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
)

func startwebserver() {
	// Grab .env stuff
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
	}
	PORT := os.Getenv("WebserverPort")

	// Start webserver
	fmt.Println("Starting Webserver on Port ", PORT)

	responseHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/main", responseHandler)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
