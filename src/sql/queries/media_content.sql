-- name: CreateMediaContent :one
INSERT INTO media_content 
    (title, description, body, content_type, created_at, user_id)
VALUES 
    ($1, $2, $3, $4, NOW(), $5)
RETURNING *;

-- name: ListMediaContent :many
SELECT * 
FROM media_content
ORDER BY id;

-- name: GetMediaContentById :one
SELECT * 
FROM media_content
WHERE id = $1;

-- name: ListMediaContentByUserId :many
SELECT * 
FROM media_content
WHERE user_id = $1
ORDER BY id;

-- name: DeleteMediaContentById :exec
DELETE FROM media_content
WHERE id = $1;
