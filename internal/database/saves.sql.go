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
    saves (created_at, updated_at, filepath)
VALUES
    (?, ?, ?) RETURNING id, created_at, updated_at, filepath
`

type CreateSaveParams struct {
	CreatedAt string
	UpdatedAt string
	Filepath  string
}

func (q *Queries) CreateSave(ctx context.Context, arg CreateSaveParams) (Safe, error) {
	row := q.db.QueryRowContext(ctx, createSave, arg.CreatedAt, arg.UpdatedAt, arg.Filepath)
	var i Safe
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Filepath,
	)
	return i, err
}
