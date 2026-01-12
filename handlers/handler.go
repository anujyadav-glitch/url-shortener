package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/storage"
)

// Request structure for shorten url
type ShortenRequest struct {
	URL string `json:"url"`
}

// Welcome Message for the home route
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Simple URL Shortener API!")
}

// handler for shorten a single url
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid JSON. Format: {\"url\":\"...\"}", http.StatusBadRequest)
		return
	}

	id := storage.SaveURL(req.URL)
	fmt.Fprintf(w, "Shortened URL: http://localhost:8080/r/%s", id)
}

// handler for redirect using shorten url
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	// Extracts the ID from the URL path (after /r/)
	id := r.URL.Path[len("/r/"):]
	original, exists := storage.GetURL(id)

	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, original, http.StatusFound)
}

// handler for shorten the urls in batch
func BatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method", http.StatusMethodNotAllowed)
		return
	}

	var urls []string
	if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
		http.Error(w, "Invalid JSON array", http.StatusBadRequest)
		return
	}

	results := make(map[string]string)
	for _, u := range urls {
		id := storage.SaveURL(u)
		results[u] = "http://localhost:8080/r/" + id
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
