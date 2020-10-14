// Code generated by sqlc. DO NOT EDIT.
// source: user_accounts.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUserAccount = `-- name: CreateUserAccount :one
INSERT INTO user_account(full_name, initials, email, username, created_at, password_hash, role_code)
  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio
`

type CreateUserAccountParams struct {
	FullName     string    `json:"full_name"`
	Initials     string    `json:"initials"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	CreatedAt    time.Time `json:"created_at"`
	PasswordHash string    `json:"password_hash"`
	RoleCode     string    `json:"role_code"`
}

func (q *Queries) CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, createUserAccount,
		arg.FullName,
		arg.Initials,
		arg.Email,
		arg.Username,
		arg.CreatedAt,
		arg.PasswordHash,
		arg.RoleCode,
	)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const deleteUserAccountByID = `-- name: DeleteUserAccountByID :exec
DELETE FROM user_account WHERE user_id = $1
`

func (q *Queries) DeleteUserAccountByID(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserAccountByID, userID)
	return err
}

const getAllUserAccounts = `-- name: GetAllUserAccounts :many
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio FROM user_account WHERE username != 'system'
`

func (q *Queries) GetAllUserAccounts(ctx context.Context) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Email,
			&i.Username,
			&i.PasswordHash,
			&i.ProfileBgColor,
			&i.FullName,
			&i.Initials,
			&i.ProfileAvatarUrl,
			&i.RoleCode,
			&i.Bio,
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

const getMemberData = `-- name: GetMemberData :many
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio FROM user_account
  WHERE username != 'system'
  AND user_id NOT IN (SELECT user_id FROM project_member WHERE project_id = $1)
`

func (q *Queries) GetMemberData(ctx context.Context, projectID uuid.UUID) ([]UserAccount, error) {
	rows, err := q.db.QueryContext(ctx, getMemberData, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAccount
	for rows.Next() {
		var i UserAccount
		if err := rows.Scan(
			&i.UserID,
			&i.CreatedAt,
			&i.Email,
			&i.Username,
			&i.PasswordHash,
			&i.ProfileBgColor,
			&i.FullName,
			&i.Initials,
			&i.ProfileAvatarUrl,
			&i.RoleCode,
			&i.Bio,
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

const getRoleForUserID = `-- name: GetRoleForUserID :one
SELECT username, role.code, role.name FROM user_account
  INNER JOIN role ON role.code = user_account.role_code
WHERE user_id = $1
`

type GetRoleForUserIDRow struct {
	Username string `json:"username"`
	Code     string `json:"code"`
	Name     string `json:"name"`
}

func (q *Queries) GetRoleForUserID(ctx context.Context, userID uuid.UUID) (GetRoleForUserIDRow, error) {
	row := q.db.QueryRowContext(ctx, getRoleForUserID, userID)
	var i GetRoleForUserIDRow
	err := row.Scan(&i.Username, &i.Code, &i.Name)
	return i, err
}

const getUserAccountByID = `-- name: GetUserAccountByID :one
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio FROM user_account WHERE user_id = $1
`

func (q *Queries) GetUserAccountByID(ctx context.Context, userID uuid.UUID) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountByID, userID)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const getUserAccountByUsername = `-- name: GetUserAccountByUsername :one
SELECT user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio FROM user_account WHERE username = $1
`

func (q *Queries) GetUserAccountByUsername(ctx context.Context, username string) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, getUserAccountByUsername, username)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const setUserPassword = `-- name: SetUserPassword :one
UPDATE user_account SET password_hash = $2 WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio
`

type SetUserPasswordParams struct {
	UserID       uuid.UUID `json:"user_id"`
	PasswordHash string    `json:"password_hash"`
}

func (q *Queries) SetUserPassword(ctx context.Context, arg SetUserPasswordParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, setUserPassword, arg.UserID, arg.PasswordHash)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const updateUserAccountInfo = `-- name: UpdateUserAccountInfo :one
UPDATE user_account SET bio = $2, full_name = $3, initials = $4, email = $5
  WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio
`

type UpdateUserAccountInfoParams struct {
	UserID   uuid.UUID `json:"user_id"`
	Bio      string    `json:"bio"`
	FullName string    `json:"full_name"`
	Initials string    `json:"initials"`
	Email    string    `json:"email"`
}

func (q *Queries) UpdateUserAccountInfo(ctx context.Context, arg UpdateUserAccountInfoParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountInfo,
		arg.UserID,
		arg.Bio,
		arg.FullName,
		arg.Initials,
		arg.Email,
	)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const updateUserAccountProfileAvatarURL = `-- name: UpdateUserAccountProfileAvatarURL :one
UPDATE user_account SET profile_avatar_url = $2 WHERE user_id = $1
  RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio
`

type UpdateUserAccountProfileAvatarURLParams struct {
	UserID           uuid.UUID      `json:"user_id"`
	ProfileAvatarUrl sql.NullString `json:"profile_avatar_url"`
}

func (q *Queries) UpdateUserAccountProfileAvatarURL(ctx context.Context, arg UpdateUserAccountProfileAvatarURLParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserAccountProfileAvatarURL, arg.UserID, arg.ProfileAvatarUrl)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE user_account SET role_code = $2 WHERE user_id = $1 RETURNING user_id, created_at, email, username, password_hash, profile_bg_color, full_name, initials, profile_avatar_url, role_code, bio
`

type UpdateUserRoleParams struct {
	UserID   uuid.UUID `json:"user_id"`
	RoleCode string    `json:"role_code"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (UserAccount, error) {
	row := q.db.QueryRowContext(ctx, updateUserRole, arg.UserID, arg.RoleCode)
	var i UserAccount
	err := row.Scan(
		&i.UserID,
		&i.CreatedAt,
		&i.Email,
		&i.Username,
		&i.PasswordHash,
		&i.ProfileBgColor,
		&i.FullName,
		&i.Initials,
		&i.ProfileAvatarUrl,
		&i.RoleCode,
		&i.Bio,
	)
	return i, err
}
