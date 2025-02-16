package main

import (
    "fmt"
    "log"
    "net/http"
    "url-shortener/handlers"
)

func main() {
    fs := http.FileServer(http.Dir("./frontend"))
    http.Handle("/", fs)  // Serves static frontend files

    http.HandleFunc("/shorten", handlers.ShortenURL) // API for shortening URLs
    http.HandleFunc("/r/", handlers.RedirectURL) // ✅ Registering `/r/` correctly

    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
