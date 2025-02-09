-- name: CreateSave :one
INSERT INTO
    saves (created_at, updated_at, system_id, filename)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetSave :one
SELECT
    *
FROM
    saves
where
    id = ?;
