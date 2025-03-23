-- +goose Up
CREATE TABLE refresh_tokens (
    token CHAR(64) PRIMARY KEY, -- use 32 bytes (256 bits) = 64 digits in hex
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    expired_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP
);

-- +goose Down
DROP TABLE refresh_tokens;
