package services

import (
	"quebrada_api/internal/domain/entities"
)

type IUserService interface {
	RegisterUser(user entities.User) error
	UpdateUser(user entities.User) error
	GetProfile(ID int64) (entities.User, error)
	DeactivateUser(ID int64) error
	CheckEmailExist(email string) error
}
