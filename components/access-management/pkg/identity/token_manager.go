package identity

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"quebrada_api/internal/domain/entities"
	"time"
)

type AuthCustomClaims struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
	//Policies []string
	jwt.RegisteredClaims
}

type TokenManager struct {
	secretKey string
	issure    string
}

func NewTokenManager(secret string, issure string) *TokenManager {
	return &TokenManager{
		secretKey: secret,
		issure:    issure,
	}
}

func (s TokenManager) GenerateToken(user entities.User) (string, error) {

	roles := make([]string, 0)
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	claims := &AuthCustomClaims{
		Name:  user.Name,
		Email: user.Email,
		Roles: roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			Issuer:    s.issure,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (s TokenManager) ReadToken(token string) (AuthCustomClaims, error) {
	claims := &AuthCustomClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(tokenJWT *jwt.Token) (interface{}, error) {
		return tokenJWT, nil
	})
	if err != nil {
		return *claims, err
	}
	if !tkn.Valid {
		return *claims, errors.New("token is invalid")
	}
	return *claims, nil
}
