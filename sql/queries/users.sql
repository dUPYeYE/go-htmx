-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = ?;


-- name: CreateUser :one
INSERT INTO users (id, name, email, password, created_at, updated_at)
VALUES (
  ?,
  ?,
  ?,
  ?,
  ?,
  ?
)
RETURNING *;

-- name: ResetPassword :one
UPDATE users
SET password = ?
WHERE id = ?
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
