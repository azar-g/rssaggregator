package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithErrorJson(w http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Println("Server error:", message)
	}
	type ErrorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, ErrorResponse{Error: message})

}

/*
*respondWithJSON sends a JSON response with the specified status code and payload.
  - It sets the Content-Type header to application/json and handles any errors that occur during JSON mar

shalling.
  - @param w http.ResponseWriter to write the response to.
  - @param status int HTTP status code to set for the response.
  - @param payload interface{} The data to be marshalled into JSON and sent in the response body.
*/
func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(data)
}
