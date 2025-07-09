package main

import (
	"fmt"
	"net/http"

	"github.com/azar-g/rssaggregator/internal/auth"
	"github.com/azar-g/rssaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	}
}
