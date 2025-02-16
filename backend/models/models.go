package models

// URLRequest represents the structure for incoming URL shortening requests.

type URLRequest struct {
    OriginalURL string `json:"original_url"`
}

// URLResponse represents the structure for shortened URL responses.
type URLResponse struct {
    ShortURL string `json:"short_url"`
}