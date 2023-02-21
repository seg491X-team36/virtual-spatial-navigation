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

-- name: GetPendingInvites :many
SELECT * FROM invites WHERE invites.user_id = $1 AND invites.experiment_id NOT IN 
(SELECT experiment_id FROM experiment_results WHERE experiment_results.user_id = $1)
ORDER BY invites.created_at ASC;