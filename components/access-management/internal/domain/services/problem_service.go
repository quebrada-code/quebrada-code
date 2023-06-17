package services

import (
	"github.com/google/uuid"
	"quebrada_api/internal/communication"
	"quebrada_api/internal/domain/domain_errors"
	"quebrada_api/internal/domain/entities"
	domain "quebrada_api/internal/domain/repositories"
	"quebrada_api/pkg/logger"
	"time"
)

type IProblemService interface {
	CreateProblem(problem entities.Problem) error
	UpdateProblem(problem entities.Problem) error
	GetProblems() ([]entities.Problem, error)
	GetProblemById(problemId uint) (entities.Problem, error)
	GetProblemsByCategory(categoryId uint) ([]entities.Problem, error)
	SubmitProblem(userId uint, problemId uint, sulutionCode string) error
}
type ProblemService struct {
	problemRepository domain.IRepository[entities.Problem]
	publisher         communication.Publisher
}

func NewProblemService(
	problemRepository domain.IRepository[entities.Problem],
	publisher communication.Publisher,
) *ProblemService {
	return &ProblemService{
		problemRepository: problemRepository,
		publisher:         publisher,
	}
}

func (p ProblemService) CreateProblem(problem entities.Problem) error {
	err := p.problemRepository.Insert(problem)
	if err != nil {
		logger.Error(err)
		return domain_errors.InsertUserError{}
	}
	return nil
}

func (p ProblemService) UpdateProblem(problem entities.Problem) error {
	return p.problemRepository.Update(problem)
}

func (p ProblemService) GetProblems() ([]entities.Problem, error) {
	return p.problemRepository.GetAll()
}

func (p ProblemService) GetProblemById(problemId uint) (entities.Problem, error) {
	return p.problemRepository.GetByID(problemId)
}

func (p ProblemService) GetProblemsByCategory(categoryId uint) ([]entities.Problem, error) {
	panic("implement me")
}

func (p ProblemService) SubmitProblem(userId uint, problemId uint, sulutionCode string) error {

	problem, err := p.GetProblemById(problemId)
	if err != nil {
		return err
	}
	event := communication.SolutionSubmittedEvent{
		UserId:       userId,
		ProblemId:    problemId,
		TestCode:     problem.TestCode,
		SolutionCode: sulutionCode,
		Language:     "python",
		CreatedAt:    time.Now(),
	}

	key := uuid.New()
	return p.publisher.Send("solution-submitted", event, key.String())

}
