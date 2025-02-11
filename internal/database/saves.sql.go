// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: saves.sql

package database

import (
	"context"
)

const createSave = `-- name: CreateSave :one
INSERT INTO
    saves (created_at, updated_at, system_id, filename)
VALUES
    (?, ?, ?, ?) RETURNING id, created_at, updated_at, system_id, filename, md5_hash
`

type CreateSaveParams struct {
	CreatedAt string
	UpdatedAt string
	SystemID  string
	Filename  string
}

func (q *Queries) CreateSave(ctx context.Context, arg CreateSaveParams) (Safe, error) {
	row := q.db.QueryRowContext(ctx, createSave,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.SystemID,
		arg.Filename,
	)
	var i Safe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SystemID,
		&i.Filename,
		&i.Md5Hash,
	)
	return i, err
}

const getSaveByID = `-- name: GetSaveByID :one
SELECT
    id, created_at, updated_at, system_id, filename, md5_hash
FROM
    saves
where
    id = ?
`

func (q *Queries) GetSaveByID(ctx context.Context, id int64) (Safe, error) {
	row := q.db.QueryRowContext(ctx, getSaveByID, id)
	var i Safe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SystemID,
		&i.Filename,
		&i.Md5Hash,
	)
	return i, err
}

const getSaves = `-- name: GetSaves :many
select
    id, created_at, updated_at, system_id, filename, md5_hash
from
    saves
LIMIT
    ?
OFFSET
    ?
`

type GetSavesParams struct {
	Limit  int64
	Offset int64
}

func (q *Queries) GetSaves(ctx context.Context, arg GetSavesParams) ([]Safe, error) {
	rows, err := q.db.QueryContext(ctx, getSaves, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Safe
	for rows.Next() {
		var i Safe
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.SystemID,
			&i.Filename,
			&i.Md5Hash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
