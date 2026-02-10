-- name: CreateFeed :one
Insert Into feeds (id, created_at, updated_at, name, url, user_id)
Values ($1, $2, $3, $4, $5, $6)
Returning *;

-- name: GetFeeds :many
Select * From feeds;