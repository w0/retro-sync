-- +goose up
CREATE TABLE saves (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    filepath TEXT NOT NULL
);
