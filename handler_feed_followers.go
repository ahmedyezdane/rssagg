package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmedyezdane/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollower(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	parameters := params{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&parameters)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while decoding json: %v", err))
		return
	}

	feed_follower, err := cfg.DB.CreateFeedFollower(r.Context(), database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    parameters.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating feed_follower : %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsFollowerToFeedsFollower(feed_follower))
}
