-- name: CreateTask :one
INSERT INTO tasks (
    id,
    owner_id,
    name,
    updated_at,
    created_at,
    date,
    priority
) VALUES ($1, $2, $3, NOW(), NOW(), $4, $5) -- skip date and time focus
RETURNING *;

-- name: GetUserTasksWeek :many
SELECT * FROM tasks
WHERE owner_id = $1 AND date >= $2 AND date <= $3
ORDER BY is_done ASC, priority DESC, created_at ASC;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;

-- name: DoneTaskByID :one
UPDATE tasks
SET is_done = true, updated_at = NOW()
WHERE id = $1 AND owner_id = $2
RETURNING *;

-- name: UndoneTaskByID :one
UPDATE tasks
SET is_done = false, updated_at = NOW()
WHERE id = $1 AND owner_id = $2
RETURNING *;

-- name: GetTaskByDate :many
SELECT * FROM tasks
WHERE owner_id = $1 AND date = $2
ORDER BY is_done ASC, priority DESC;

-- name: UpdateTaskKeys :one
UPDATE tasks
SET keys = $1, updated_at = NOW()
WHERE owner_id = $2 AND id = $3
RETURNING *;

-- name: DeleteTaskByID :exec
DELETE FROM tasks
WHERE id = $1 AND owner_id = $2;
