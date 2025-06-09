package main

import (
	"time"

	"github.com/htooanttko/rssagg_api/internal/database"
)

type User struct {
	ID        int32	`json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
}

func dbUsertoUser(user database.User) User {
	return User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name: user.Name,
		ApiKey: user.ApiKey,
	}
}