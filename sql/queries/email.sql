-- name: ListEmails :many
SELECT e.*
FROM emails e
ORDER BY e.created_at DESC;

-- name: CreateEmail :exec
INSERT INTO emails (id, recipient, subject, body, priority)
VALUES ($1, $2, $3, $4, $5);
