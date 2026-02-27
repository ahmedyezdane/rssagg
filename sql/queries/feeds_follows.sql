-- name: CreateFeedFollow :one
Insert Into feeds_follows(id,created_at,updated_at,user_id,feed_id)
Values ($1, $2, $3, $4, $5)
Returning *;

-- name: GetFeedFollowsOfUser :many
Select * From feeds_follows Where user_id = $1;

-- name: DeleteFeedFollows :exec
Delete From feeds_follows Where user_id = $1 And id = $2;