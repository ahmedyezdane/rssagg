-- name: CreateFeedFollow :one
Insert Into feeds_follows(id,created_at,updated_at,user_id,feed_id)
Values ($1, $2, $3, $4, $5)
Returning *;