package main

import (
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {

	// Respond with a generic error message
	respondWithErrorJson(w, http.StatusBadRequest, "Something went wrong, please try again later.")
}
