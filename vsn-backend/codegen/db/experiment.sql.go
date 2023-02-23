// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: experiment.sql

package db

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

const createExperiment = `-- name: CreateExperiment :one
INSERT INTO
    experiments (
        name,
        description,
        config
    )
VALUES ($1, $2, $3)
RETURNING id, name, description, config
`

type CreateExperimentParams struct {
	Name        string
	Description string
	Config      json.RawMessage
}

func (q *Queries) CreateExperiment(ctx context.Context, arg CreateExperimentParams) (Experiment, error) {
	row := q.db.QueryRowContext(ctx, createExperiment, arg.Name, arg.Description, arg.Config)
	var i Experiment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Config,
	)
	return i, err
}

const getExperiment = `-- name: GetExperiment :one
SELECT id, name, description, config
FROM experiments
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetExperiment(ctx context.Context, id uuid.UUID) (Experiment, error) {
	row := q.db.QueryRowContext(ctx, getExperiment, id)
	var i Experiment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Config,
	)
	return i, err
}

const getExperiments = `-- name: GetExperiments :many
SELECT id, name, description, config
FROM experiments
`

func (q *Queries) GetExperiments(ctx context.Context) ([]Experiment, error) {
	rows, err := q.db.QueryContext(ctx, getExperiments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Experiment
	for rows.Next() {
		var i Experiment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Config,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}