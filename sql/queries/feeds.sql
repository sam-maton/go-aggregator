-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE name = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name as user_name FROM feeds
LEFT JOIN users ON feeds.user_id = users.id;

-- name: DeleteFeeds :exec
DELETE FROM feeds;