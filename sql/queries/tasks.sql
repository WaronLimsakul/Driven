-- name: CreateTask :one
INSERT INTO tasks (
    id,
    owner_id,
    name,
    date,
    priority
) VALUES ($1, $2, $3, $4, $5) -- skip date and time focus
RETURNING *;

-- name: GetUserTasksWeek :many
SELECT * FROM tasks
WHERE owner_id = $1 AND date >= $2 AND date <= $3
ORDER BY priority DESC;

-- name: GetTaskByID :one
SELECT * FROM tasks
WHERE id = $1;
