-- name: Count :one
SELECT COUNT(*)
FROM users;

-- name: FindOne :one
SELECT *
FROM users
WHERE (
    sqlc.narg(id)::uuid IS NULL
    OR id = sqlc.narg(id)::uuid
  )
  AND (
    sqlc.narg(email)::varchar IS NULL
    OR email = sqlc.narg(email)::varchar
  )
LIMIT 1;

-- name: FindPage :many
SELECT *
FROM users
LIMIT $1 OFFSET $2;

-- name: Create :one
INSERT INTO users (
    first_name,
    last_name,
    email,
    password_hash,
    password_salt
  )
VALUES (
    sqlc.arg(first_name),
    sqlc.arg(last_name),
    sqlc.arg(email),
    sqlc.arg(password_hash),
    sqlc.arg(password_salt)
  )
RETURNING *;

-- name: Update :one
UPDATE users
SET first_name = COALESCE(sqlc.narg(first_name)::varchar(255), first_name),
  last_name = COALESCE(sqlc.narg(last_name)::varchar(255), last_name),
  password_hash = COALESCE(
    sqlc.narg(password_hash)::varchar(255),
    password_hash
  ),
  password_salt = COALESCE(
    sqlc.narg(password_salt)::varchar(255),
    password_salt
  )
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: Delete :one
DELETE FROM users
WHERE id = sqlc.arg(id)
RETURNING *;
