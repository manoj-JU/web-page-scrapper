package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type DBConf struct {
	userName, password, dbName, host, port string
}

var dbConf DBConf

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

//initialized DBConf object with environment values
func initializeDBConf() DBConf {
	return DBConf{
		userName: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		dbName:   os.Getenv("DB_DB_NAME"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
	}
}

// Main function
func main() {
	fmt.Println("db-persistor service started")
	dbConf = initializeDBConf()

	// Init router
	r := mux.NewRouter()
	r.HandleFunc("/amazon/products", CreateAmazonProductHandler).Methods("POST")
	r.HandleFunc("/amazon/products", GetAmazonProductsHanlder).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
