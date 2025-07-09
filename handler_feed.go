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
func (apiCfg *apiConfig) createUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		Name string `json:"name" required:"true"`
		Url  string `json:"url" required:"true"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	})

	if err != nil {
		respondWithErrorJson(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed: %v", err))
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

// func (apiCfg *apiConfig) getUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {

// }
