package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	return string(hashpassword), nil
}

func CheckPassword(password string, hashpassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err
}
