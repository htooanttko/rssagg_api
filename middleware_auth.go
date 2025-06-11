package main

import (
	"net/http"

	"github.com/htooanttko/rssagg_api/internal/auth"
	"github.com/htooanttko/rssagg_api/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey,err := auth.GetAPIKey(r.Header)
	if err != nil {
		resWithErr(w, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		resWithErr(w, http.StatusNotFound,err.Error())
	}
	handler(w, r, user)
	}
}