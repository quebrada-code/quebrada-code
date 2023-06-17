package migrations

import (
	"gorm.io/gorm"
	"quebrada_api/internal/domain/entities"
	"quebrada_api/pkg/logger"
)

func Migrate(database *gorm.DB) {

	logger.Info("Start migrations")

	err := database.AutoMigrate(&entities.User{}, &entities.Role{}, &entities.Policy{}, &entities.UserAccess{})
	if err != nil {
		logger.Error(err)
		panic("Failed on migrate database")
	}

	database.Model(&entities.User{})
	database.Model(&entities.Role{})
	database.Model(&entities.Policy{})
	database.Model(&entities.UserAccess{})

}
