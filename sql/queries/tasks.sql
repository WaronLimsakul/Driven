-- name: CreateTask :one
INSERT INTO tasks (
    id,
    owner_id,
    name,
    date,
    priority
) VALUES ($1, $2, $3, $4, $5) -- skip date and time focus
RETURNING *;
