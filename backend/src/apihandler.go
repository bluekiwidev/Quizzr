package main

import (
	"fmt"
	"log"
	"net/http"
)

func startwebserver(PORT string) {

	// Start webserver
	fmt.Println("\n Starting Webserver on Port ", PORT)

	http.HandleFunc("/checkusername", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		usernametocheck := r.URL.Query().Get("username")
		fmt.Println("API HANDLER: Received username to check availibility:", usernametocheck)
		usernamevalidcheck(usernametocheck)

		if usernamevalidcheck(usernametocheck) == 409 {
			w.WriteHeader(http.StatusConflict)
		} else if usernamevalidcheck(usernametocheck) == 200 {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(PORT, nil))
}
