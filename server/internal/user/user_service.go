package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Akshat-z/Chat-app/util"
	"github.com/golang-jwt/jwt/v4"
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

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("input password should be correct %w", err)
	}
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

type MyJWTClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const secretKey = "secretKey"

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	userData, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, (fmt.Errorf("user didn't exit %w", err))
	}
	err = util.CheckPassword(req.Password, userData.Password)
	if err != nil {
		return nil, (fmt.Errorf("password is incorrect %w", err))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		Email:    req.Email,
		Username: userData.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(userData.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &LoginUserRes{
		ID:          strconv.Itoa(int(userData.ID)),
		UserName:    userData.UserName,
		accessToken: ss,
	}, nil
}
