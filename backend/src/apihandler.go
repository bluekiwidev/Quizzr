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
		usernametocheck := r.URL.Query().Get("username")
		fmt.Println("API HANDLER: Received username to check availibility: ", usernametocheck)
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
