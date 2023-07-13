package main

import (
	"fmt"
	"log"
	"net/http"
)

var shortURLs map[string]string

func main() {
	shortURLs = make(map[string]string)

	http.HandleFunc("/", createShortURL)
	http.HandleFunc("/{id}", getFullURL)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.URL.Path != "/" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	fullURL := string(body)

	id := generateID()
	shortURLs[id] = fullURL

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", id)
}

func getFullURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	id := r.URL.Path[1:]
	fullURL, ok := shortURLs[id]
	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Location", fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func generateID() string {
	// generate a unique id for short URL
	// implementation is omitted for brevity
	return "EwHXdJfB"
}
