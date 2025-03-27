-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name TEXT UNIQUE NOT NULL,
    keys TEXT,
    date DATE NOT NULL,
    priority INT NOT NULL,
    time_focus INT NOT NULL DEFAULT 0 -- time focus in minutes
);

-- +goose Down
DROP TABLE tasks;
