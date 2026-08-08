package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func appenduser(username string, email string, uncryptedpword string) bool {
	fmt.Println(username)
	fmt.Println(email)
	fmt.Println(uncryptedpword)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(uncryptedpword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("FAILED TO HASH PWORD...")
	}

	password := string(hashedBytes)
	fmt.Println(password)

	fmt.Println(compareemail(email))
	if compareemail(email) == false {
		if adduser(username, email, password) == true {
			return (true)
		} else {
			return (false)
		}
	} else {
		return (false)
	}
}

func login(email string, password string) (bool, string) {
	start := time.Now()

	// Check if email exists
	if compareemail(email) != true {
		elapsed := time.Since(start)
		if elapsed < 3*time.Second {
			time.Sleep(3*time.Second - elapsed)
		}
		fmt.Println("[LOG] Email does not exist")
		return false, "0"
	}

	// Check if password is correct
	if comparepassword(email, password) != true {
		elapsed := time.Since(start)
		if elapsed < 3*time.Second {
			time.Sleep(3*time.Second - elapsed)
		}
		fmt.Println("[LOG] Password is incorrect")
		return false, "0"
	}

	// Generare a random session key
	keyBytes := make([]byte, 32)
	rand.Read(keyBytes)

	// Store the session in Redis
	storesession(base64.URLEncoding.EncodeToString(keyBytes), email)
	fmt.Println("[LOG] Session stored in Redis for email:", email)

	// Return the session key
	return true, base64.URLEncoding.EncodeToString(keyBytes)
}
