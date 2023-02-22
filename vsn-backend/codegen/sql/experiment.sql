-- name: CreateExperiment :one
INSERT INTO
    experiments (
        name,
        description,
        config
    )
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetExperiment :one
SELECT *
FROM experiments
WHERE id = $1
LIMIT 1;

-- name: GetExperiments :many
SELECT *
FROM experiments;