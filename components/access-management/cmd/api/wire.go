package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
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

func InitProblemController(db *gorm.DB, kafkaProducer *kafka.Producer) controllers.ProblemController {
	repository := ioc.ProvideProblemRepository(db)
	service := ioc.ProvideProblemService(repository, kafkaProducer)
	problemController := ioc.ProvideProblemController(service)
	return *problemController

}

func initAuthController(db *gorm.DB) controllers.AuthController {
	wire.Build(ioc.ProvideAuthRepositoryRepostiory, ioc.ProvideAuthService, ioc.ProvideAuthController)

	return controllers.AuthController{}
}
