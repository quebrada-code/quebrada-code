package main

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"quebrada_api/internal/app/controllers"
	"quebrada_api/internal/config"
	"quebrada_api/ioc"
	"quebrada_api/pkg/identity"
)

func InitAuthController(db *gorm.DB, config config.STMPConfig) controllers.AuthController {
	repository := ioc.ProvideAuthRepositoryRepostiory(db)
	passwordHash := ioc.ProvidePasswordHash("PASSWORD")
	emailSender := ioc.ProviderEmailSender(config)
	authService := ioc.ProvideAuthService(repository, passwordHash, db, emailSender)
	authController := ioc.ProvideAuthController(authService, identity.TokenManager{})
	return *authController
}

func initAuthController(db *gorm.DB) controllers.AuthController {
	wire.Build(ioc.ProvideAuthRepositoryRepostiory, ioc.ProvideAuthService, ioc.ProvideAuthController)

	return controllers.AuthController{}
}
