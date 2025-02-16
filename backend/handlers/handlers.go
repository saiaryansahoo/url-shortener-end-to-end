package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/models"
	"url-shortener/utils"
)

var urlStore = make(map[string]string)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req models.URLRequest
	json.NewDecoder(r.Body).Decode(&req)
	shortURL := utils.GenerateShortURL(req.OriginalURL)
	urlStore[shortURL] = req.OriginalURL

	resp := models.URLResponse{ShortURL: shortURL}
	json.NewEncoder(w).Encode(resp)
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path

    // Ensure the request is at least "/r/{shortURL}"
    if len(path) <= len("/r/") {
        http.Error(w, "Invalid short URL", http.StatusBadRequest)
        return
    }

    shortURL := path[len("/r/"):] // Extract the short URL safely

    if originalURL, ok := urlStore[shortURL]; ok {
        http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
    } else {
        http.Error(w, "URL not found", http.StatusNotFound)
    }
}
