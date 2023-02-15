-- name: GetInvite :one
SELECT *
FROM invites
WHERE id = $1
LIMIT 1;

-- name: GetInvitesByExperimentId :many
SELECT *
FROM invites
WHERE supervised = $1
AND experiment_id = $2;

-- name: CreateInvite :one
INSERT INTO
    invites (
        user_id,
        experiment_id,
        supervised
    )
VALUES
    ($1, $2, $3)
RETURNING *;