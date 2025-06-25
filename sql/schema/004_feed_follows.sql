-- +goose Up
CREATE TABLE feed_follows (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id SERIAL NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id,feed_id),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feed_follows;