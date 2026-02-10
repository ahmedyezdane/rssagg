package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmedyezdane/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing json: %v", err))
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Creating feed: %v", err))
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while fetching feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}
