-- name: CreateUser :one
Insert Into users (id,created_at,updated_at,name)
Values ($1, $2, $3, $4)
Returning *;

-- name: GetUserByAPIKey :one
Select * From users Where api_key = $1;