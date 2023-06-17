package ioc

import (
	"gorm.io/gorm"
	"quebrada_api/internal/domain/entities"
	domain "quebrada_api/internal/domain/repositories"
	"quebrada_api/internal/infrastructure/repositories"
)

func ProvideAuthRepositoryRepostiory(DB *gorm.DB) domain.IRepository[entities.User] {
	return repositories.Repository[entities.User]{DB: DB}
}

func ProvideProblemRepository(DB *gorm.DB) domain.IRepository[entities.Problem] {
	return repositories.Repository[entities.Problem]{DB: DB}
}
