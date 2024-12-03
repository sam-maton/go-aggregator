-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)
SELECT new_feed_follow.*, feeds.name as feed_name, users.name as user_name
FROM new_feed_follow
INNER JOIN feeds ON new_feed_follow.feed_id = feeds.id
INNER JOIN users ON new_feed_follow.user_id = users.id;