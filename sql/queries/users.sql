-- name: CreateUser :one
INSERT INTO users(name, api_key, created_at,updated_at) 
VALUES (
    $1, 
    encode(sha256(random()::text::bytea), 'hex'),
    $2, 
    $3
)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users WHERE api_key = $1;