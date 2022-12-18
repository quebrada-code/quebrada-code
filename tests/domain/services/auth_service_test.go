package services

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"quebrada_api/internal/domain/entities"
	"quebrada_api/internal/domain/services"
	"testing"
)

type AuthServiceSuite struct {
	suite.Suite
	conn    *sql.DB
	DB      *gorm.DB
	mock    sqlmock.Sqlmock
	service services.UserService
}

//
//func (suite *AuthServiceSuite) SetupSuite() {
//	var (
//		err error
//	)
//	suite.conn, suite.mock, err = sqlmock.New()
//	assert.NoError(suite.T(), err)
//
//	suite.DB, err = gorm.Open(sqlite.Open("file:test?mode=memory&cache=shared&_fk=1"), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Silent),
//	})
//	assert.NoError(suite.T(), err)
//
//	err = suite.DB.AutoMigrate(&entities.User{})
//	assert.NoError(suite.T(), err)
//
//	repo := repositories.Repository[entities.User]{DB: suite.DB}
//
//	hasher :=
//
//	suite.service = services.NewUserService(repo, nil,nil, nil)
//
//	suite.DB.Create(&entities.User{
//		Name:     "Fulano Ciclano",
//		Password: "q1w2e3r4",
//		Email:    "fulano@gmail",
//		Active:   true,
//	})
//}

func (suite *AuthServiceSuite) TestCheckEmail() {
	err := suite.service.CheckEmailExist("fulano@gmail")
	suite.NotNil(err)
	suite.Equal(err.Error(), "existe usuário cadastrado com esse e-mail")
}

// TestRegisterUserWithUserExist A test case.
func (suite *AuthServiceSuite) TestRegisterUserWithUserExist() {
	user := entities.User{
		Name:     "Fulano Ciclano",
		Password: "q1w2e3r4",
		Email:    "fulano@gmail",
		Active:   true,
	}
	err := suite.service.RegisterUser(user)
	suite.NotNil(err)
	suite.Equal(err.Error(), "existe usuário cadastrado com esse e-mail")
}

func TestGenericRepositorySuite(t *testing.T) {
	suite.Run(t, new(AuthServiceSuite))
}
