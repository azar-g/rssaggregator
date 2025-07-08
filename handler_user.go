package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azar-g/rssaggregator/internal/auth"
	"github.com/azar-g/rssaggregator/internal/database"
	"github.com/google/uuid"
)

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

func (apiCfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)

	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, fmt.Sprintf("Unauthorized: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

	if err != nil {
		respondWithErrorJson(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
