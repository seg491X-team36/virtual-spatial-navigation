// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO
    users (
        email,
        state,
        source
    )
VALUES
    ($1, $2, $3)
RETURNING id, email, state, source
`

type CreateUserParams struct {
	Email  string
	State  UserAccountState
	Source string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.State, arg.Source)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.State,
		&i.Source,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, state, source
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.State,
		&i.Source,
	)
	return i, err
}

const getUsersByState = `-- name: GetUsersByState :many
SELECT id, email, state, source
FROM users
WHERE state = $1
`

func (q *Queries) GetUsersByState(ctx context.Context, state UserAccountState) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsersByState, state)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.State,
			&i.Source,
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

const updateUserState = `-- name: UpdateUserState :many
UPDATE 
    users
SET
    state=$1
WHERE
    id in ($2::UUID[])
RETURNING id, email, state, source
`

type UpdateUserStateParams struct {
	State   UserAccountState
	UserIds []uuid.UUID
}

func (q *Queries) UpdateUserState(ctx context.Context, arg UpdateUserStateParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, updateUserState, arg.State, pq.Array(arg.UserIds))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.State,
			&i.Source,
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
