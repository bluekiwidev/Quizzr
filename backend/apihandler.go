package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func startwebserver() {
	fmt.Println("Starting Webserver on Port 8080")

	responseHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/main", responseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
