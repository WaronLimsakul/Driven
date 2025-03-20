-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    email,
    hashed_password,
    created_at,
    updated_at
) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
