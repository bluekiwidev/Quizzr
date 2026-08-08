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

func login(email string, password string) bool {
	start := time.Now()

	// Check if email exists
	if compareemail(email) != true {
		elapsed := time.Since(start)
		if elapsed < 3*time.Second {
			time.Sleep(3*time.Second - elapsed)
		}
		return false
	}

	// Check if password is correct
	if comparepassword(email, password) != true {
		elapsed := time.Since(start)
		if elapsed < 3*time.Second {
			time.Sleep(3*time.Second - elapsed)
		}
		return false
	}
	return true
}
