// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user_schema.users (username, password_hash, salt, email)
VALUES ($1, $2, $3, $4) RETURNING id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at
`

type CreateUserParams struct {
	Username     string `db:"username" json:"username"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	Salt         string `db:"salt" json:"salt"`
	Email        string `db:"email" json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.PasswordHash,
		arg.Salt,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Salt,
		&i.Email,
		&i.EmailVerified,
		&i.FirstName,
		&i.LastName,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at FROM user_schema.users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.PasswordHash,
			&i.Salt,
			&i.Email,
			&i.EmailVerified,
			&i.FirstName,
			&i.LastName,
			&i.DateOfBirth,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at FROM user_schema.users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Salt,
		&i.Email,
		&i.EmailVerified,
		&i.FirstName,
		&i.LastName,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE user_schema.users
SET first_name = $2 AND last_name = $3 AND updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, username, password_hash, salt, email, email_verified, first_name, last_name, date_of_birth, created_at, updated_at
`

type UpdateUserParams struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FirstName *string   `db:"first_name" json:"first_name"`
	LastName  *string   `db:"last_name" json:"last_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.ID, arg.FirstName, arg.LastName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.PasswordHash,
		&i.Salt,
		&i.Email,
		&i.EmailVerified,
		&i.FirstName,
		&i.LastName,
		&i.DateOfBirth,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}