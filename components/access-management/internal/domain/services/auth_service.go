package services

import (
	"errors"
	"github.com/xlzd/gotp"
	"gorm.io/gorm"
	"quebrada_api/internal/domain/domain_errors"
	"quebrada_api/internal/domain/entities"
	domain "quebrada_api/internal/domain/repositories"
	"quebrada_api/internal/domain/spec"
	"quebrada_api/internal/infrastructure/senders"
	"quebrada_api/pkg/identity"
	"quebrada_api/pkg/logger"
)

type IAuthService interface {
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	DeleteUser(userId uint) error
	SignInWithEmail(email, password string) (entities.User, error)
	GeneratePasswordResetToken(email string) (string, error)
	ResetPassword(email string, token string, newPassword string) error
	GenerateEmailConfirmationToken(user entities.User) error
	ConfirmEmail(userId uint, token string) error
	//AddToRole(userId uint, roleName string) error
	//RemoveToRole(userId uint, roleName string) error
	//GrantPolicy(userId uint, policyName string) error
	//RevokePolicy(userId uint, policyName string) error
}

type VerificationCode struct {
	Name             string
	VerificationCode string
}

type AuthService struct {
	userRepository domain.IRepository[entities.User]
	passwordHasher identity.IPasswordHasher
	sender         senders.ISender
	DB             *gorm.DB
}

func NewAuthService(
	userRepository domain.IRepository[entities.User],
	passwordHasher identity.IPasswordHasher,
	sender senders.ISender,
	DB *gorm.DB,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		sender:         sender,
		DB:             DB,
	}
}

func (a *AuthService) CreateUser(user entities.User) error {
	exist, err := a.CheckEmailExist(user.Email)
	if err != nil || exist {
		logger.Error(err)
		return domain_errors.UserAlreadyExistsError{}
	}

	passwordHash, err := a.passwordHasher.HashPassword(user.Password)
	if err != nil {
		logger.Error(err)
		return domain_errors.HashPasswordError{}
	}
	user.PasswordHash = passwordHash

	verificationCode := gotp.RandomSecret(6)
	user.VerificationCode = verificationCode

	err = a.userRepository.Insert(user)
	if err != nil {
		logger.Error(err)
		return domain_errors.InsertUserError{}
	}

	go func() {
		err := a.GenerateEmailConfirmationToken(user)
		if err != nil {
			logger.Error("failed to send welcome email.")
		}
	}()

	return nil
}

func (a *AuthService) UpdateUser(user entities.User) error {
	return a.userRepository.Update(user)
}

func (a *AuthService) DeleteUser(userId uint) error {
	return a.userRepository.Delete(userId)
}

func (a *AuthService) SignInWithEmail(email, password string) (entities.User, error) {
	user, err := a.FindByEmail(email)
	if err != nil {
		return entities.User{}, err
	}

	passwordHash, err := a.passwordHasher.HashPassword(password)
	if err != nil {
		logger.Error(err)
		return entities.User{}, errors.New("nenhum usuário cadastrado")
	}

	if user.PasswordHash != passwordHash {
		return entities.User{}, errors.New("senha é inválida")
	}

	return user, nil
}

func (a *AuthService) FindByEmail(email string) (entities.User, error) {
	result, err := a.userRepository.Query(spec.GetUserWithEmailSpec(email))
	if err != nil {
		return entities.User{}, err
	}
	if len(result) == 0 {
		return entities.User{}, errors.New("nao foi encontrado nenhum usuário")
	} else {
		return result[0], nil
	}
}

func (a *AuthService) GeneratePasswordResetToken(email string) (string, error) {

	user, err := a.FindByEmail(email)
	if err != nil {
		return "", err
	}

	user.ResetToken = gotp.RandomSecret(6)

	a.DB.Save(user)

	return user.ResetToken, nil
}

func (a *AuthService) ResetPassword(email string, token string, newPassword string) error {
	user, err := a.FindByEmail(email)
	if err != nil {
		return err
	}

	if token == user.ResetToken {
		password, err := a.passwordHasher.HashPassword(newPassword)
		if err != nil {
			return err
		}
		user.PasswordHash = password
		a.DB.Save(user)
		return nil
	} else {
		return errors.New("reset token é inválido")
	}
}

func (a *AuthService) GenerateEmailConfirmationToken(user entities.User) error {
	subject := "Quebrada Code - Seja bem-vindo"
	template := "/Users/marcos.lopes/projects/pessoal/plataform/quebrada_api/resources/welcome.html"
	data := VerificationCode{Name: user.Name, VerificationCode: user.VerificationCode}
	err := a.sender.Send([]string{user.Email}, subject, template, data)
	if err != nil {
		logger.Error("failed to send welcome email.")
	}
	return err
}

func (a *AuthService) ConfirmEmail(userId uint, token string) error {
	var user entities.User
	tx := a.DB.First(&user, user)
	if tx.Error != nil {
		return tx.Error
	}
	if user.VerificationCode != token {
		return errors.New("verification code invalid")
	}

	user.EmailConfirmed = true
	user.VerificationCode = token
	tx = a.DB.Save(&user)
	return tx.Error
}

//
//func (a *AuthService) AddToRole(userId uint, roleName string) error {
//	panic("implement me")
//}
//
//func (a *AuthService) RemoveToRole(userId uint, roleName string) error {
//	panic("implement me")
//}
//
//func (a *AuthService) GrantPolicy(userId uint, policyName string) error {
//	panic("implement me")
//}
//
//func (a *AuthService) RevokePolicy(userId uint, policyName string) error {
//	panic("implement me")
//}

func (a *AuthService) CheckEmailExist(email string) (bool, error) {
	query, err := a.userRepository.Query(spec.GetUserWithEmailSpec(email))
	if err != nil {
		return false, err
	}
	return len(query) > 0, nil
}
