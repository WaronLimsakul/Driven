// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    email,
    hashed_password,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id, name, email, hashed_password, created_at, updated_at, focus_time
`

type CreateUserParams struct {
	ID             uuid.UUID
	Name           string
	Email          string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FocusTime,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, name, email, hashed_password, created_at, updated_at, focus_time FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FocusTime,
	)
	return i, err
}

const resetUsers = `-- name: ResetUsers :exec
DELETE FROM users
`

func (q *Queries) ResetUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetUsers)
	return err
}
