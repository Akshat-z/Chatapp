package user

import (
	"context"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repo Repository) Service {
	return &service{
		repo,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// TODO: hash password
	hashPassword := req.Password
	u := &User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: hashPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &CreateUserRes{
		UserName: r.UserName,
		ID:       strconv.Itoa(int(r.ID)),
		Email:    r.Email,
	}, nil
}
