-- name: CreateGame :one
INSERT INTO
    games (created_at, updated_at, system_id, filename)
VALUES
    (?, ?, ?, ?) RETURNING *;
