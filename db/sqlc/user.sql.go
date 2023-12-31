// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email,
  hashed_password,
  username
) VALUES (
  $1, $2, $3
) RETURNING id, email, hashed_password, username, is_email_verified, password_changed_at, created_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	Username       string `json:"username"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.Username,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, hashed_password, username, is_email_verified, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.Username,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE($1, hashed_password),
  password_changed_at = COALESCE($2, password_changed_at),
  username = COALESCE($3, username),
  email = COALESCE($4, email),
  is_email_verified = COALESCE($5, is_email_verified)
WHERE
  username = $3
RETURNING id, email, hashed_password, username, is_email_verified, password_changed_at, created_at
`

type UpdateUserParams struct {
	HashedPassword    sql.NullString `json:"hashed_password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	Username          sql.NullString `json:"username"`
	Email             sql.NullString `json:"email"`
	IsEmailVerified   sql.NullBool   `json:"is_email_verified"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.HashedPassword,
		arg.PasswordChangedAt,
		arg.Username,
		arg.Email,
		arg.IsEmailVerified,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.Username,
		&i.IsEmailVerified,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
