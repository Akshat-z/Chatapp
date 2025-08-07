package user

import (
	"context"
	"database/sql"
)

type IDBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db IDBTX
}

func New(db IDBTX) Repository {
	return &repository{db: db}
}

const (
	queryInsertUser = "INSERT INTO users (username, email, password) VALUES ($1,$2,$3) returning id"
)

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	err := r.db.QueryRowContext(ctx, queryInsertUser, user.UserName, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return &User{}, err
	}
	user.ID = int64(lastInsertId)
	return user, nil
}
