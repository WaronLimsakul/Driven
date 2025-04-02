-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    keys TEXT, -- postgres TEXT support multi-line text, so it works fine with textarea
    date DATE NOT NULL,
    priority INT NOT NULL,
    is_done BOOL NOT NULL DEFAULT false,
    time_focus INT NOT NULL DEFAULT 0 -- time focus in minutes
);

-- +goose Down
DROP TABLE tasks;
