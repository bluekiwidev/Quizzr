package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// LATER ME, ADD BCRYPT IF IT WORKS FINE
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

	fmt.Println(doesemailexist(email))
	if doesemailexist(email) == false {
		if adduser(username, email, password) == true {
			return (true)
		} else {
			return (false)
		}
	} else {
		return (false)
	}
}

func login(email string, password string, username string) bool {
	start := time.Now()

	if doesemailexist(email) == true {
		return true
	} else {
		return false
	}

	elapsed := time.Since(start)
	if elapsed < 3*time.Second {
		time.Sleep(3*time.Second - elapsed)
	}
}
