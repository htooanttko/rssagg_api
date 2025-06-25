package main

import (
	"net/http"
	"strconv"

	"github.com/htooanttko/rssagg_api/internal/database"
)

func (apiCfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r * http.Request, user database.User) {
	limitStr := r.URL.Query().Get("limit")
	limit := 10
	
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})

	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	resWithJSON(w, http.StatusOK, dbPostsToPosts(posts))
}