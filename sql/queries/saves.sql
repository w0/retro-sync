-- name: CreateSave :one
INSERT INTO
    saves (created_at, updated_at, filepath)
VALUES
    (?, ?, ?) RETURNING *;
