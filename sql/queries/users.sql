-- name: CreateUser :one
INSERT INTO users (
    id, created, modified, name, api_key
) VALUES ( $1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;


-- name: GetUserByAPIKey :one
SELECT * FROM users where api_key=$1;

-- name: GetUserByID :one
SELECT * FROM users where id=$1;

