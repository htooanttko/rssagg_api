package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/htooanttko/rssagg_api/internal/database"
)

func (apiCfg apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		FeedId int32 `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	
	err := decoder.Decode(&params)
	if err != nil {
		resWithErr(w, http.StatusInternalServerError,err.Error())
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: params.FeedId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	resWithJSON(w, http.StatusCreated,dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg apiConfig) handlerFeedFollowGet(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollows, err := apiCfg.DB.GetFeedFollowByUser(r.Context(),user.ID)
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	resWithJSON(w, http.StatusOK,dbFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDstr := chi.URLParam(r, "FeedFollowID")
	feedFollowID64,err := strconv.ParseInt(feedFollowIDstr, 10, 32) // in strconv, sec param is base int (10 for decimal, 16 for hexadecimal)
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedFollowID := int32(feedFollowID64)

	err = apiCfg.DB.DeleteFeedFollowByUser(r.Context(),database.DeleteFeedFollowByUserParams{
		UserID: user.ID,
		FeedID: feedFollowID,
	})
	if err != nil {
		resWithErr(w, http.StatusInternalServerError,err.Error())
		return
	}

	resWithJSON(w, http.StatusOK, struct{}{})
}