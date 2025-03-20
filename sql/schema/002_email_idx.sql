-- +goose Up
-- default B-tree
CREATE INDEX user_email_idx ON users (email);

-- +goose Down
DROP INDEX user_email_idx;
