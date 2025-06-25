-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollowByUser :many
SELECT * FROM feed_follows WHERE user_id = $1;

-- name: DeleteFeedFollowByUser :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;