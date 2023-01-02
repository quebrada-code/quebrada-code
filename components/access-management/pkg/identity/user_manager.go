package identity

//
//import (
//	"errors"
//	"github.com/xlzd/gotp"
//	"gorm.io/gorm"
//	"quebrada_api/pkg/logger"
//)
//
//type IUserManager interface {
//	CreateUser(user User) error
//	UpdateUser(user User) error
//	DeleteUser(userId uint) error
//	FindByUserName(user User) error
//	SignInWithEmail(email, password string) (User, error)
//	GeneratePasswordResetToken(userId uint)
//	ResetPassword(userId string, token string, newPassword string) Result
//	GenerateEmailConfirmationToken(user User) error
//	ConfirmEmail(userId uint, token string) error
//	AddToRole(userId uint, roleName string) error
//	RemoveToRole(userId uint, roleName string) error
//}
//
//type UserManager struct {
//	db           *gorm.DB
//	hasher       IPasswordHasher
//	emailService ISender
//}
//
//func NewUserManager(db *gorm.DB, passwordHasher IPasswordHasher) *UserManager {
//	return &UserManager{
//		db:     db,
//		hasher: passwordHasher,
//	}
//}
//
//func (u *UserManager) CreateUser(user User) error {
//
//	if err := u.CheckEmailExist(user.Email); err != nil {
//		return err
//	}
//
//	passwordHash, err := u.hasher.HashPassword(user.Password)
//	if err != nil {
//		return err
//	}
//
//	user.Password = passwordHash
//
//	// hash password
//	tx := u.db.Create(user)
//	if tx.Error != nil {
//		return errors.New("failed insert user")
//	}
//
//	go u.GenerateEmailConfirmationToken(user)
//
//	go func() {
//		subject := "Quebrada Code - Seja bem-vindo"
//		template := "welcome.html"
//		err := u.emailService.Send([]string{user.Email}, subject, template)
//		if err != nil {
//			logger.Error("failed to send welcome email.")
//		}
//	}()
//
//	return nil
//
//}
//
//func (u *UserManager) UpdateUser(user User) error {
//	tx := u.db.Updates(user)
//	if tx.Error != nil {
//		logger.Error("Failed update user")
//	}
//	return tx.Error
//}
//
//func (u *UserManager) DeleteUser(userId uint) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) FindByUserName(user User) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) SignInWithEmail(email, password string) (User, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) GeneratePasswordResetToken(userId uint) {
//
//	panic("implement me")
//}
//
//func (u *UserManager) ResetPassword(userId string, token string, newPassword string) Result {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) GenerateEmailConfirmationToken(user User) {
//	u.db.First(&user, user.ID)
//
//	verificationCode := gotp.RandomSecret(6)
//
//	user.VerificationCode = verificationCode
//	u.db.Save(&user)
//
//	subject := "Quebrada Code - Seja Bem Vindo"
//	template := "verification_code.html"
//	go func() {
//		err := u.emailService.Send([]string{user.Email}, subject, template)
//		if err != nil {
//			logger.Error("failed to send account verification email.")
//		}
//	}()
//}
//
//func (u *UserManager) ConfirmEmail(userId uint, token string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) AddToRole(userId uint, roleName string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) RemoveToRole(userId uint, roleName string) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (u *UserManager) CheckEmailExist(email string) error {
//	return nil
//}
