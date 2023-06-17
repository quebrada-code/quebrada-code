package identity

import (
	"github.com/stretchr/testify/assert"
	"quebrada_api/internal/domain/entities"
	"quebrada_api/pkg/identity"
	"testing"
)

func TestGenerateTokenWithSuccess(t *testing.T) {
	t.Run("Generate token with sucess", func(t *testing.T) {
		tokenManager := identity.NewTokenManager("secret", "Bikash")
		user := entities.User{
			ID:             0,
			Name:           "Marcos",
			Email:          "marcos.mota",
			EmailConfirmed: true,
			PasswordHash:   "q1w2e3r",
			Roles: []entities.Role{
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

func TestReadTokenWithSucess(t *testing.T) {
	t.Run("Read token with sucess", func(t *testing.T) {
		tokenManager := identity.NewTokenManager("secret", "teste")
		user := entities.User{
			ID:             0,
			Name:           "Marcos",
			Email:          "marcos.mota",
			EmailConfirmed: true,
			PasswordHash:   "q1w2e3r",
			Roles: []entities.Role{
				{
					ID:   1,
					Name: "Admin",
				},
			},
			Policies: nil,
		}

		token, _ := tokenManager.GenerateToken(user)
		claims, _ := tokenManager.ReadToken(token)

		assert.NotNil(t, claims)
		//assert.Nil(t, err)
	})
}
