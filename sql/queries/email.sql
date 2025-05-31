-- name: ListEmails :many
SELECT e.*, es.status
FROM emails e
         INNER JOIN public.email_status es ON e.id = es.email_id
         INNER JOIN (SELECT MAX(id) id
                     FROM email_status
                     GROUP BY email_id) st ON es.id = st.id
ORDER BY e.created_at DESC;

-- name: CreateEmail :one
INSERT INTO emails (recipient, subject, body, priority)
VALUES ($1, $2, $3, $4)
RETURNING id;
