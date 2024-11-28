-- name: AddRefreshToken :one
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES (
  ?,
  ?,
  ?
)
RETURNING *;

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token = ?;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = ?, updated_at = ?
WHERE token = ?;
