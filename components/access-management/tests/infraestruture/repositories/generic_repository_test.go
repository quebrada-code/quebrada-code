package repositories

import (
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm/logger"
	domain "quebrada_api/internal/domain/repositories"
	"quebrada_api/internal/domain/spec"
	repositories "quebrada_api/internal/infrastructure/repositories"
	"testing"
)

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type EntityTest struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Password string
	Email    string
	Active   bool
}

func GetUsersActivedSpec() spec.Specification {
	return spec.And(spec.Equal("active", true))
}

type RepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	repo domain.IRepository[EntityTest]
}

// this function executes before the test suite begins execution
func (suite *RepositorySuite) SetupSuite() {
	var (
		err error
	)
	suite.conn, suite.mock, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.DB, err = gorm.Open(sqlite.Open("file:test?mode=memory&cache=shared&_fk=1"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(suite.T(), err)

	err = suite.DB.AutoMigrate(&EntityTest{})
	assert.NoError(suite.T(), err)

	suite.repo = repositories.Repository[EntityTest]{DB: suite.DB}
	assert.IsType(suite.T(), repositories.Repository[EntityTest]{}, suite.repo)

	for i := 0; i < 5; i++ {
		suite.DB.Create(&EntityTest{
			Name:     faker.NAME,
			Password: faker.PASSWORD,
			Email:    faker.Email,
			Active:   true,
		})
	}

	for i := 0; i < 5; i++ {
		suite.DB.Create(&EntityTest{
			Name:     faker.FirstName,
			Password: faker.FirstNameFemale,
			Email:    faker.Email,
			Active:   false,
		})
	}
}

func (suite *RepositorySuite) TestGetAll() {

	result, err := suite.repo.GetAll()
	suite.Nil(err)
	suite.Equal(10, len(result))
}

func (suite *RepositorySuite) TestQuery() {

	result, err := suite.repo.Query(GetUsersActivedSpec())
	suite.Nil(err)
	suite.Equal(5, len(result))
}

func (suite *RepositorySuite) TestGetById() {

	result, err := suite.repo.GetByID(1)
	suite.Nil(err)
	suite.NotNil(result)
}

func (suite *RepositorySuite) TestInsert() {
	user := EntityTest{
		Name:     "user_teste",
		Password: "q1w2e3r4",
		Email:    "q1w2e3r4",
		Active:   false,
	}
	err := suite.repo.Insert(user)
	suite.Nil(err)
}

func TestGenericRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
