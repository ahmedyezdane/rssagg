-- name: CreateFeedFollower :one
Insert Into feeds_followers(id,created_at,updated_at,user_id,feed_id)
Values ($1, $2, $3, $4, $5)
Returning *;