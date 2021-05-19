package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
}

func GetCustomError(err error) *CustomError {
	return &CustomError{
		Message: err.Error(),
	}
}

func HandleError(err error, statusCode int, w http.ResponseWriter) {
	customError := GetCustomError(err)
	log.Printf("error occured: %v", *customError)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(*customError); err != nil {
		log.Printf("error occured: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}
