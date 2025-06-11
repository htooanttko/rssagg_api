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

type Feed struct {
	ID        int32 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    int32 `json:"user_id"`
}

func dbFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
		UserID:    feed.UserID,
	}
}

func dbFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed,len(feeds))
	for idx, feed := range feeds {
		result[idx] = dbFeedToFeed(feed)
	}
	return result 
} 