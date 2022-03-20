package main

import (
	"fmt"
	"net/http"

	// use third party package for handling complex routing requests
	"github.com/gorilla/mux"
)

func main() {
	// Define router to handle complex routing requests
	r := mux.NewRouter()

	// Process dynamic requests: Process incoming requests from users who browse the website, log into their account
	// Request contains all information needed for processing requests (r.URL.Query().Get("token"), (r.FormValue("email")))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello World!")
	})
	// Use gorilla/mux instead of regular http package to handle complex routing requests
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Println("title: ", title, "page: ", page)
		// get the book
		// navigate to page
	})

	// Serve static assets: Serve Javascript, CSS and images to browsers
	// 1. Build a fileserver
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Accept connections: Start the server on port 8080
	http.ListenAndServe(":8090", r)
}
