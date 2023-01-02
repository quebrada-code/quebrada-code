package repositories

import "quebrada_api/internal/domain/spec"

type IRepository[T interface{}] interface {
	GetAll() ([]T, error)
	GetByID(ID interface{}) (T, error)
	Insert(entity T) error
	Update(entity T) error
	Delete(ID interface{}) error
	Query(specification spec.Specification) ([]T, error)
}
