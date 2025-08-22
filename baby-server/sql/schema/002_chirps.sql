-- +goose up
CREATE TABLE chirps(
    id UUID PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body TEXT NOT NULL
);





-- +goose down
DROP TABLE chirps;