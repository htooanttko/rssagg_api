package main

import (
	"database/sql"
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

type FeedFollow struct {
	ID        int32	`json:"id"`
	UserID    int32	`json:"user_id"`
	FeedID    int32	`json:"feed_id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

func dbFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func dbFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow,len(feedFollows))
	for idx, feedFollow := range feedFollows {
		result[idx] = dbFeedFollowToFeedFollow(feedFollow)
	}
	return result
}

type Post struct {
	ID         	int32	   `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      int32	   `json:"feed_id"`
}

func dbPostToPost(post database.Post) Post {
	return Post{
			ID:          post.ID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Title:       post.Title,
			Url:         post.Url,
			Description: nullStringToStringPtr(post.Description),
			PublishedAt: nullTimeToTimePtr(post.PublishedAt),
			FeedID:      post.FeedID,
	}
}

func dbPostsToPosts(posts []database.Post) []Post {
	result := make([]Post,len(posts))
	for idx, post := range posts {
		result[idx] = dbPostToPost(post)
	}
	return result
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

