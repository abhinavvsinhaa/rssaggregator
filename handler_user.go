package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abhinavvsinhaa/rssaggregator/internal/auth"
	"github.com/abhinavvsinhaa/rssaggregator/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error creating user: %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, 200, user)
}
