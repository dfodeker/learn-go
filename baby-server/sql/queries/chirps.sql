-- name: CreateChirp :one
insert into chirps (id, created_at, updated_at, user_id, body)
values (gen_random_uuid(), now(), now(), $1, $2)
RETURNING *;


-- name: GetChirpByDate :many
SELECT * FROM chirps ORDER BY created_at ASC;


-- name: GetChirpByID :one
SELECT * FROM chirps WHERE id = $1;