-- name: CreateSave :one
INSERT INTO
    saves (created_at, updated_at, system_id, filename)
VALUES
    (?, ?, ?, ?) RETURNING *;

-- name: GetSaveByID :one
SELECT
    *
FROM
    saves
where
    id = ?;

-- name: GetSaves :many
select
    *
from
    saves
LIMIT
    ?
OFFSET
    ?;
