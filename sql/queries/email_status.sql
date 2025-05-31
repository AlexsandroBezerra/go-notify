-- name: GetByEmailId :many
SELECT *
FROM email_status
WHERE status = ($1)
ORDER BY created_at;

-- name: GetLastByEmailId :one
SELECT *
FROM email_status
WHERE status = ($1)
LIMIT 1;

-- name: CreateEmailStatus :exec
INSERT INTO email_status (email_id, status)
VALUES ($1, $2);
