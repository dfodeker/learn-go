-- name: CreateUser :one
insert into users (id, created_at, updated_at, email,hashed_password)
values (gen_random_uuid(),now(), now(), $1, $2)
RETURNING *;



-- name: DeleteAllUsers :exec
DELETE FROM users;


-- name: getAllUsers :many
SELECT * FROM users ORDER BY created_at ASC;



-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;