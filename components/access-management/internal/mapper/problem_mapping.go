package mapper

import (
	"quebrada_api/internal/app/models"
	"quebrada_api/internal/domain/entities"
)

func ToProblem(model models.CreateProblem) entities.Problem {
	return entities.Problem{
		Name:        model.Name,
		ProblemCode: model.ProblemCode,
		TestCode:    model.TestCode,
		Order:       model.Order,
		Description: model.Description,
		Point:       model.Point,
		Level:       model.Level,
	}
}
