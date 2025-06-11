package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/htooanttko/rssagg_api/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	decoders := json.NewDecoder(r.Body)
	params := parameters{}
	
	err := decoders.Decode(&params)
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		resWithErr(w, http.StatusInternalServerError,err.Error())
		return
	}

	resWithJSON(w, http.StatusCreated, dbFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerFeedGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	resWithJSON(w, http.StatusOK, dbFeedsToFeeds(feeds))
}