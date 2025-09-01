-- token: the primary key - it's just a string
-- created_at
-- updated_at
-- user_id: foreign key that deletes the row if the user is deleted
-- expires_at: the timestamp when the token expires
-- revoked_at:

-- name: CreateRefreshToken :one
insert into refresh_tokens (token, user_id, expires_at, revoked_at)
values ($1, $2, $3, NULL)
RETURNING *;

-- name: GetRefreshTokenByToken :one
SELECT * FROM refresh_tokens WHERE token = $1;

-- name: RevokeRefreshToken :one
UPDATE refresh_tokens
SET revoked_at = NOW(), updated_at = NOW()
WHERE token = $1
RETURNING *;

