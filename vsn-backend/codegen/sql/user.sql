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

-- name: UpdateUserState :exec
UPDATE 
    users
SET
    state=$2
WHERE
    id=$1
RETURNING *;