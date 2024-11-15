-- name: CreateUser :one
INSERT INTO users 
    (first_name, last_name, email, password)
VALUES 
    ($1, $2, $3, $4)
RETURNING *;

-- name: ListUsers :many
SELECT * 
FROM users
ORDER BY id;

-- name: GetUserById :one
SELECT * 
FROM users
WHERE id = $1;

-- name: UpdateUserById :one
UPDATE users
SET 
    first_name = $2,
    last_name = $3,
    email = $4,
    password = $5
WHERE 
    id = $1
RETURNING *;

-- name: DeleteUserById :exec
DELETE FROM users
WHERE id = $1;
