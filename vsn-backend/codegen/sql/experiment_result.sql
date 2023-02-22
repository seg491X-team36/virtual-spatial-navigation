-- name: CreateExperimentResult :one
INSERT INTO
    experiment_results (
        id,
        user_id,
        experiment_id
    )
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetExperimentResultsByExperimentId :many
SELECT *
FROM experiment_results
WHERE experiment_id = $1;

-- name: GetExperimentResultsByUserId :many
SELECT *
FROM experiment_results
WHERE user_id = $1;