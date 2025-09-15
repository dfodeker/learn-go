-- +goose up
alter TABLE users
    ADD COLUMN is_chirpy_red BOOLEAN NOT NULL DEFAULT FALSE;


-- +goose down
alter TABLE users
    DROP COLUMN is_chirpy_red;