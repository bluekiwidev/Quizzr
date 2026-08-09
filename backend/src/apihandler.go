package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

	type Payload struct {
		Email    string `json:"Email"`
		Username string `json:"Username"`
		Password string `json:"Password"`
	}

	http.HandleFunc("/submitsignup", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fmt.Println("Received request")

		var payload Payload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if appenduser(payload.Username, payload.Email, payload.Password) == true {
			fmt.Println("IT MADE THE ACC")
			w.WriteHeader(http.StatusCreated)
		}

	})

	http.HandleFunc("/submitsignin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println("Received request")

		// Decode payload
		var payload Payload

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Check login info
		loginStatus, sessionKey := login(payload.Email, payload.Password)

		// Validate login and set cookie
		if loginStatus == true && sessionKey != "0" {
			fmt.Println("Login is a match!")
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sessionKey,
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(24 * time.Hour),
			})
			w.WriteHeader(http.StatusOK)
		} else {
			// Invalid login, return unauthorized status
			print("Email no exist or password is a nono, no login")
			w.WriteHeader(http.StatusUnauthorized)
		}

	})

	http.HandleFunc("/submitsignout", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		fmt.Println("Received request")

		// Decode payload
		token, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

	})

	log.Fatal(http.ListenAndServe(PORT, nil))
}
