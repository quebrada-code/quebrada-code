package ioc

import (
	"quebrada_api/internal/app/controllers"
	"quebrada_api/internal/domain/services"
	"quebrada_api/pkg/identity"
)

func ProvideAuthController(
	authService services.IAuthService,
	tokenManager identity.TokenManager) *controllers.AuthController {
	return controllers.NewAuthController(authService, tokenManager)
}
