-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUsersByState :many
SELECT *
FROM users
WHERE state = $1;

-- name: CreateUser :one
INSERT INTO
    users (
        email,
        state,
        source
    )
VALUES
    ($1, $2, $3)
RETURNING *;

-- name: UpdateUserState :many
UPDATE 
    users
SET
    state=$1
WHERE
    id in (sqlc.arg(user_ids)::UUID[])
RETURNING *;