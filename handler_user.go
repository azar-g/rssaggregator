package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azar-g/rssaggregator/internal/database"
	"github.com/google/uuid"
)

/**
 * createUserHandler handles the creation of a new user.
 * It decodes the request body to extract the user's name, creates a new user in the database,
 * and responds with the created user's details in JSON format.
 * If the request body is invalid or if there is an error during user creation, it responds with an appropriate error message.
 */
func (apiCfg *apiConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type Params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:      params.Name,
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

/**
 * getUser retrieves the user associated with the API key from the request header.
 * It responds with the user details in JSON format or an error message if the user is not found or if there is an internal server error.
 */
func (apiCfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
