package main

import (
	"fmt"
	"net/http"
	"url-shortener/internal/handlers"
)

func main() {
	// home
	http.HandleFunc("/", handlers.RootHandler)

	// route for a shorten a single url
	http.HandleFunc("/shorten", handlers.ShortenHandler)

	// route for redirect
	http.HandleFunc("/r/", handlers.RedirectHandler)

	// route for shorten urls in batch
	http.HandleFunc("/batch", handlers.BatchHandler)

	fmt.Println("Server is running on http://localhost:8080")
	// Start server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
