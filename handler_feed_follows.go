package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/azar-g/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) createUserFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type Params struct {
		FeedID uuid.UUID `json:"feed_id" required:"true"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Params{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErrorJson(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateUserFeedFollow(r.Context(), database.CreateUserFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithErrorJson(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create feed follow: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *apiConfig) getAllFeedFollowsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetAllFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		respondWithErrorJson(w, http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve feeds: %v", err))
		return
	}

	formattedFeedFollows := []FeedFollow{}

	for _, feedFollow := range feedFollows {

		formattedFeedFollows = append(formattedFeedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	respondWithJSON(w, http.StatusOK, formattedFeedFollows)
}

func (apiCfg *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

}

// func (apiCfg *apiConfig) getUserFeed(w http.ResponseWriter, r *http.Request, user database.User) {

// }
