// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tasks.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
    id,
    owner_id,
    name,
    date,
    priority
) VALUES ($1, $2, $3, $4, $5) -- skip date and time focus
RETURNING id, owner_id, name, keys, date, priority, isdone, time_focus
`

type CreateTaskParams struct {
	ID       uuid.UUID
	OwnerID  uuid.UUID
	Name     string
	Date     time.Time
	Priority int32
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.OwnerID,
		arg.Name,
		arg.Date,
		arg.Priority,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.Isdone,
		&i.TimeFocus,
	)
	return i, err
}

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, owner_id, name, keys, date, priority, isdone, time_focus FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.Isdone,
		&i.TimeFocus,
	)
	return i, err
}

const getUserTasksWeek = `-- name: GetUserTasksWeek :many
SELECT id, owner_id, name, keys, date, priority, isdone, time_focus FROM tasks
WHERE owner_id = $1 AND date >= $2 AND date <= $3
ORDER BY priority DESC
`

type GetUserTasksWeekParams struct {
	OwnerID uuid.UUID
	Date    time.Time
	Date_2  time.Time
}

func (q *Queries) GetUserTasksWeek(ctx context.Context, arg GetUserTasksWeekParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getUserTasksWeek, arg.OwnerID, arg.Date, arg.Date_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.Name,
			&i.Keys,
			&i.Date,
			&i.Priority,
			&i.Isdone,
			&i.TimeFocus,
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
