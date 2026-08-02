package main

import (
	"fmt"
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
		log.Fatal("\n Couldnt find .env file in backend dir. HINT: Is the .env actually in the backend dir?")
	}
	PORT := os.Getenv("WebserverPort")

	// Start webserver
	fmt.Println("\n Starting Webserver on Port ", PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(PORT, nil))
}
