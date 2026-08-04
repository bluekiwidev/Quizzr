package main

import (
	"fmt"

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
