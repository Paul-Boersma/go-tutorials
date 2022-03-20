package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handle a request for localhost:8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World")
	})

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
