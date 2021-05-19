package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AmazonProduct struct
type AmazonProduct struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	ImageURL     string `json:"image_url"`
	Description  string `json:"description"`
	Price        string `json:"price"`
	TotalReviews string `json:"total_reviews"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type URL struct {
	URL string `json:"url"`
}

func main() {
	fmt.Println("amazon-scrapper service started")

	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/scrape/amazon", ScrapeAmazonProductPageHandler).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
