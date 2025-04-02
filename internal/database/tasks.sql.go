// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: tasks.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
    id,
    owner_id,
    name,
    updated_at,
    created_at,
    date,
    priority
) VALUES ($1, $2, $3, NOW(), NOW(), $4, $5) -- skip date and time focus
RETURNING id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus
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
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.IsDone,
		&i.TimeFocus,
	)
	return i, err
}

const doneTaskByID = `-- name: DoneTaskByID :one
UPDATE tasks
SET is_done = true, updated_at = NOW()
WHERE id = $1 AND owner_id = $2
RETURNING id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus
`

type DoneTaskByIDParams struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
}

func (q *Queries) DoneTaskByID(ctx context.Context, arg DoneTaskByIDParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, doneTaskByID, arg.ID, arg.OwnerID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.IsDone,
		&i.TimeFocus,
	)
	return i, err
}

const getTaskByDate = `-- name: GetTaskByDate :many
SELECT id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus FROM tasks
WHERE owner_id = $1 AND date = $2
ORDER BY is_done ASC, priority DESC
`

type GetTaskByDateParams struct {
	OwnerID uuid.UUID
	Date    time.Time
}

func (q *Queries) GetTaskByDate(ctx context.Context, arg GetTaskByDateParams) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTaskByDate, arg.OwnerID, arg.Date)
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
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Keys,
			&i.Date,
			&i.Priority,
			&i.IsDone,
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

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus FROM tasks
WHERE id = $1
`

func (q *Queries) GetTaskByID(ctx context.Context, id uuid.UUID) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.IsDone,
		&i.TimeFocus,
	)
	return i, err
}

const getUserTasksWeek = `-- name: GetUserTasksWeek :many
SELECT id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus FROM tasks
WHERE owner_id = $1 AND date >= $2 AND date <= $3
ORDER BY is_done ASC, priority DESC, created_at ASC
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
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.Keys,
			&i.Date,
			&i.Priority,
			&i.IsDone,
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

const undoneTaskByID = `-- name: UndoneTaskByID :one
UPDATE tasks
SET is_done = false, updated_at = NOW()
WHERE id = $1 AND owner_id = $2
RETURNING id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus
`

type UndoneTaskByIDParams struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
}

func (q *Queries) UndoneTaskByID(ctx context.Context, arg UndoneTaskByIDParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, undoneTaskByID, arg.ID, arg.OwnerID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.IsDone,
		&i.TimeFocus,
	)
	return i, err
}

const updateTaskKeys = `-- name: UpdateTaskKeys :one
UPDATE tasks
SET keys = $1, updated_at = NOW()
WHERE owner_id = $2 AND id = $3
RETURNING id, owner_id, name, updated_at, created_at, keys, date, priority, is_done, time_focus
`

type UpdateTaskKeysParams struct {
	Keys    sql.NullString
	OwnerID uuid.UUID
	ID      uuid.UUID
}

func (q *Queries) UpdateTaskKeys(ctx context.Context, arg UpdateTaskKeysParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskKeys, arg.Keys, arg.OwnerID, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.Keys,
		&i.Date,
		&i.Priority,
		&i.IsDone,
		&i.TimeFocus,
	)
	return i, err
}
