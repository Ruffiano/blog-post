-- name: CreateSession :one
INSERT INTO sessions (
  id,
  user_id,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = $1 LIMIT 1;


-- name: UpdateSession :one
UPDATE sessions
SET
  refresh_token = COALESCE(sqlc.narg(refresh_token), refresh_token),
  user_agent = COALESCE(sqlc.narg(user_agent), user_agent),
  client_ip = COALESCE(sqlc.narg(client_ip), client_ip),
  expires_at = COALESCE(sqlc.narg(expires_at), expires_at)
WHERE
  user_id = sqlc.arg(user_id)
RETURNING *;
