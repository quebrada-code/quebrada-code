package repositories

import (
	"gorm.io/gorm"
	"quebrada_api/internal/domain/spec"
)

type Repository[T interface{}] struct {
	DB *gorm.DB
}

func (r Repository[T]) GetAll() ([]T, error) {
	limit, offset := -1, -1
	var models []T

	err := r.DB.Limit(limit).Offset(offset).Find(&models).Error

	if err != nil {
		return nil, err
	}

	result := make([]T, 0, len(models))
	for _, row := range models {
		result = append(result, row)
	}

	return result, nil
}

func (r Repository[T]) GetByID(ID interface{}) (T, error) {
	var model T
	tx := r.DB.First(&model, ID)
	if tx.Error != nil {
		return model, tx.Error
	}

	return model, tx.Error

}

func (r Repository[T]) Insert(entity T) error {
	tx := r.DB.Create(&entity)
	return tx.Error
}

func (r Repository[T]) Update(entity T) error {
	panic("implement me")
}

func (r Repository[T]) Delete(ID interface{}) error {
	var model T
	tx := r.DB.Delete(&model, ID)

	return tx.Error
}

func (r Repository[M]) Query(specification spec.Specification) ([]M, error) {
	// retreive reords by some criteria
	var models []M
	err := r.DB.Where(specification.GetQuery(), specification.GetValues()...).Find(&models).Error
	// handle error

	if err != nil {
		return nil, err
	}

	// mapp all records into Entities
	result := make([]M, 0, len(models))
	for _, row := range models {
		result = append(result, row)
	}

	return result, nil
}
