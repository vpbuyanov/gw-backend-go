package common

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("create hashed password was failed: %v", err.Error())
	}

	return string(hashPassword), nil
}

func CompareHashAndPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}
