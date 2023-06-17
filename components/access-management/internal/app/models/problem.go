package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ProblemModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       string `json:"level"`
	Point       int    `json:"point"`
	TestCode    string `json:"testCode"`
	ProblemCode string `json:"problemCode"`
	Order       int16  `json:"order"`
}

type CreateProblem struct {
	*ProblemModel
}

func (a CreateProblem) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.Description, validation.Required),
		validation.Field(&a.ProblemCode, validation.Required),
		validation.Field(&a.TestCode, validation.Required),
	)
}

type ResponseProblem struct {
	*ProblemModel
	Id uint `json:"id"`
}

type SubimtProblemModel struct {
	UserId       uint   `json:"userId"`
	ProblemId    uint   `json:"problemId"`
	SolutionCode string `json:"solutionCode"`
}

func (a SubimtProblemModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.ProblemId, validation.Required),
		validation.Field(&a.SolutionCode, validation.Required),
	)
}
