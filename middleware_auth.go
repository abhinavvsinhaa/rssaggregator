package main

import (
	"fmt"
	"net/http"

	"github.com/abhinavvsinhaa/rssaggregator/internal/auth"
	"github.com/abhinavvsinhaa/rssaggregator/internal/database"
)

func (apiCfg *apiConfig) authMiddleware(handler func(http.ResponseWriter, *http.Request, database.User)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.ExtractAPIKeyFromHeader(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Authentication error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not fetch user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
