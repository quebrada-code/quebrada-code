package mapper

import (
	"quebrada_api/internal/app/models"
	"quebrada_api/internal/domain/entities"
)

func ToEntity(model models.SignUpModel) entities.User {
	return entities.User{
		Name:     model.Name,
		Password: model.Password,
		Email:    model.Email,
		Active:   true,
	}
}
