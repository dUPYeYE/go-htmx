-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;


-- name: CreateUser :one
INSERT INTO users (name, email, password)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: ResetPassword :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
