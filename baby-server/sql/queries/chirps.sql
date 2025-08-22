-- name: CreateChirp :one
insert into chirps (id, created_at, updated_at, user_id, body)
values (gen_random_uuid(), now(), now(), $1, $2)
RETURNING *;


