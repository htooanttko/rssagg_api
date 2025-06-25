package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/htooanttko/rssagg_api/internal/database"
)

func (apiCfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	
	err := decoder.Decode(&params)
	if err != nil {
		resWithErr(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		resWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	resWithJSON(w, http.StatusCreated, dbUsertoUser(user))
}

func (apiCfg *apiConfig) handlerUserGet(w http.ResponseWriter, r *http.Request, user database.User) {
	resWithJSON(w, http.StatusOK, dbUsertoUser(user))
}