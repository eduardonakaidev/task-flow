package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// hashPassword gera o hash da senha utilizando bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}
	return string(bytes), nil
}
