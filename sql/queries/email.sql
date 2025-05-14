-- name: ListEmails :many
SELECT *
FROM emails
ORDER BY emails.created_at DESC;

-- name: CreateEmail :one
INSERT INTO emails (recipient, subject, body, priority, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;