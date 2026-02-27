package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ahmedyezdane/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
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

	feed_follow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    parameters.FeedId,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating feed_follow : %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsFollowToFeedsFollow(feed_follow))
}

func (cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	user_feed_follows, err := cfg.DB.GetFeedFollowsOfUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting user feed_follows : %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedsFollowsToFeedsFollows(user_feed_follows))
}

func (cfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing feedFollowId: %v", err))
		return
	}

	err = cfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		ID:     feedFollowId,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while deleting user feed_follows : %v", err))
		return
	}

	respondWithJSON(w, 200, feedFollowId)
}
