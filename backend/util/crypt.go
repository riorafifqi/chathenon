package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateBcrypt(password string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("bcrypt error") // TODO: Change to constant
	}

	passHash := string(hash)

	return &passHash, nil
}

func CompareHashPassword(hashPassword []byte, userPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, userPassword)
	if err != nil {
		return bcrypt.ErrMismatchedHashAndPassword
	}
	return nil
}
