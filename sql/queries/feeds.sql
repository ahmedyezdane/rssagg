-- name: CreateFeed :one
Insert Into feeds (id, created_at, updated_at, name, url, user_id)
Values ($1, $2, $3, $4, $5, $6)
Returning *;

-- name: GetFeeds :many
Select * From feeds;

-- name: GetNextFeedsToFetch :many
SELECT * from feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;