package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash hashes the provided password using bcrypt.
func GeneratePasswordHash(password string) (string, error) {
	bytePass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswordHash compares the provided password with the hashed password.
func ComparePasswordHash(hashedPassword string, password string) error {
	byteHashPass := []byte(hashedPassword)
	bytePass := []byte(password)

	err := bcrypt.CompareHashAndPassword(byteHashPass, bytePass)
	return err
}
