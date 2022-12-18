package services

import (
	"errors"
	"github.com/xlzd/gotp"
	"quebrada_api/internal/domain/entities"
	"quebrada_api/internal/domain/repositories"
	"quebrada_api/internal/domain/spec"
	"quebrada_api/internal/infrastructure/services"
	"quebrada_api/pkg/hash"
)

type UserService struct {
	userRepository repositories.IRepository[entities.User]
	emaiService    services.IEmailService
	hasher         hash.PasswordHasher
}

func NewUserService(
	repository repositories.IRepository[entities.User],
	hasher hash.PasswordHasher,
	emailService services.IEmailService) UserService {
	return UserService{
		repository,
		emailService,
		hasher,
	}
}

func (s *UserService) RegisterUser(user entities.User) error {
	if err := s.CheckEmailExist(user.Email); err != nil {
		return err
	}

	passwordHash, err := s.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = passwordHash

	// hash password
	if err := s.userRepository.Insert(user); err != nil {
		return err
	}

	//TODO: Enviar e-mail de validação
	verificationCode := gotp.RandomSecret(6)

	go s.SendUserVerificationEmail(
		user.Email,
		user.Name,
		verificationCode)
	return nil
}

func (s *UserService) CheckEmailExist(email string) error {
	result, err := s.userRepository.Query(spec.GetUserWithEmailSpec(email))
	if err != nil {
		return err
	}

	if len(result) > 0 {
		return errors.New("existe usuário cadastrado com esse e-mail")
	}
	return nil
}

func (s *UserService) SendUserVerificationEmail(email, name, code string) {
	s.emaiService.SendEmail([]string{email}, "", "")
}
