package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func getShortenURL(url string) string {
	url = "http://short.url"
	return url
}

func shortenURLHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
	urlBytes, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Bad request", http.StatusBadRequest)
	}

	url := string(urlBytes)

	shortURL := getShortenURL(url)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(shortURL))
	if err != nil {
		http.Error(res, "No way!", http.StatusBadRequest)
	}

}

func originalURLHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "No GET-method", http.StatusBadRequest)
	}

	shortURL2 := req.URL.Path[:1]

	originalURL := getOriginalURL(shortURL2)

	if originalURL == "" {
		http.Error(res, "No URL", http.StatusBadRequest)
	}

	res.Header().Set("Location", originalURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}

//func checkID(res http.ResponseWriter, req *http.Request) {
//	vars = mux.Vars(req)
//	id, ok := vars["id"]
//}

func getOriginalURL(shortURL2 string) string {
	shortURL2 = "https://google.com"
	return shortURL2
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/{id}`, originalURLHandler)
	r.HandleFunc(`/`, shortenURLHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
