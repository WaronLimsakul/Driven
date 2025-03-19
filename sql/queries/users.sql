-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    hashed_password,
    created_at,
    updated_at
) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING *;
