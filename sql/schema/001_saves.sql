-- +goose Up
CREATE TABLE saves (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    system_id TEXT NOT NULL,
    filename TEXT NOT NULL,
    md5_hash VARCHAR(32)
);

-- +goose Down
DROP TABLE saves;
