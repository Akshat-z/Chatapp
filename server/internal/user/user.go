package user

import (
	"context"
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserReq struct {
	UserName string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRes struct {
	ID       string `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id" db:"id"`
	UserName    string `json:"username" db:"username"`
}

type Repository interface {
	CreateUser(context.Context, *User) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
}

type Service interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	Login(context.Context, *LoginUserReq) (*LoginUserRes, error)
}
