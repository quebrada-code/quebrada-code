package identity

import (
	"github.com/stretchr/testify/assert"
	"quebrada_api/pkg/identity"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	t.Run("Generate token with sucess", func(t *testing.T) {
		tokenManager := identity.NewTokenManager("secret", "Bikash")
		user := identity.User{
			ID:             0,
			Name:           "Marcos",
			Email:          "marcos.mota",
			EmailConfirmed: true,
			PasswordHash:   "q1w2e3r",
			Roles: []identity.Role{
				{
					ID:   1,
					Name: "Admin",
				},
			},
			Policies: nil,
		}
		token, err := tokenManager.GenerateToken(user)
		assert.Nil(t, err)
		assert.NotNil(t, token)
	})
}

func TestReadToken(t *testing.T) {
	t.Run("Read token with sucess", func(t *testing.T) {
		tokenManager := identity.NewTokenManager("secret", "teste")
		user := identity.User{
			ID:             0,
			Name:           "Marcos",
			Email:          "marcos.mota",
			EmailConfirmed: true,
			PasswordHash:   "q1w2e3r",
			Roles: []identity.Role{
				{
					ID:   1,
					Name: "Admin",
				},
			},
			Policies: nil,
		}

		token, err := tokenManager.GenerateToken(user)
		claims, err := tokenManager.ReadToken(token)

		assert.NotNil(t, claims)
		assert.Nil(t, err)
	})
}
