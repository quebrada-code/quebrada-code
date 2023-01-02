package identity

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

type IPasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyHashedPassword(password string, hashed string) error
}

type PasswordHasher struct {
	salt string
}

func NewPasswordHasher(salt string) *PasswordHasher {
	return &PasswordHasher{salt: salt}
}

func (h *PasswordHasher) HashPassword(password string) (string, error) {
	hash := sha1.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}

func (h *PasswordHasher) VerifyHashedPassword(password string, hashed string) error {
	hashPassword, err := h.HashPassword(password)
	if err != nil {
		return err
	}
	if hashPassword != hashed {
		return errors.New("passwprd is invalid")
	}
	return nil
}
