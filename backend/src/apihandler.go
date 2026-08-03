package main

import (
	"fmt"
	"log"
	"net/http"
)

func startwebserver(PORT string) {

	// Start webserver
	fmt.Println("\n Starting Webserver on Port ", PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Fatal(http.ListenAndServe(PORT, nil))
}
