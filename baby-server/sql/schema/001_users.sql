-- +goose up
CREATE TABLE users(
     id UUID PRIMARY KEY NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);




-- +goose down
DROP TABLE users;