package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"quebrada_api/internal/app/controllers"
	"quebrada_api/internal/app/models"
	"quebrada_api/internal/app/router"
	"quebrada_api/internal/config"
	"quebrada_api/ioc"
	"quebrada_api/migrations"
	"quebrada_api/pkg/identity"
	"testing"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx
}

func InitAuthController(
	db *gorm.DB,
	config config.STMPConfig,
) controllers.AuthController {
	repository := ioc.ProvideAuthRepositoryRepostiory(db)
	passwordHash := ioc.ProvidePasswordHash("PASSWORD")
	emailSender := ioc.ProviderEmailSender(config)
	authService := ioc.ProvideAuthService(repository, passwordHash, db, emailSender)
	authController := ioc.ProvideAuthController(authService, identity.TokenManager{})
	return *authController
}

type AuthControllerSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *AuthControllerSuite) SetupSuite() {
	var (
		err error
	)
	_, _, err = sqlmock.New()
	assert.NoError(suite.T(), err)

	db, err := gorm.Open(sqlite.Open("file:test?mode=memory&cache=shared&_fk=1"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(suite.T(), err)

	migrations.Migrate(db)
	suite.router = SetUpRouter()

	smptConfig := config.STMPConfig{
		Host:     "server.com",
		Port:     545,
		User:     "user",
		Password: "password",
	}

	routerManager := router.Router{
		AuthController: InitAuthController(db, smptConfig),
	}

	api := suite.router.Group("/api")
	{
		routerManager.Init(api)
	}
}

func (suite *AuthControllerSuite) TestShouldValidateModelWithSuccess() {

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	type signUpModelRequest struct {
		Name            string
		Email           string
		Nickname        string
		ZipCode         string
		Password        string
		ConfirmPassword string
	}

	var tests = []struct {
		name    string
		resquet *signUpModelRequest
	}{
		{"All field empty", &signUpModelRequest{
			Name:            "",
			Email:           "",
			Nickname:        "",
			ZipCode:         "",
			Password:        "",
			ConfirmPassword: "",
		}},
		{"Zip code incorret", &signUpModelRequest{
			Name:            "Marcos Mota",
			Email:           "marcos.mota287@gmail.com",
			Nickname:        "mota287",
			ZipCode:         "14058085",
			Password:        "q1w2e3r4",
			ConfirmPassword: "q1w2e3r4",
		}},
		{"Senhas não correspondem", &signUpModelRequest{
			Name:            "Marcos Mota",
			Email:           "marcos.mota287@gmail.com",
			Nickname:        "mota287",
			ZipCode:         "14058-085",
			Password:        "q1w2e",
			ConfirmPassword: "q1w2e3r4",
		}},
	}

	for _, test := range tests {
		buf, _ := json.Marshal(test.resquet)

		req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/auth/sign-up", bytes.NewReader(buf))
		suite.router.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		var res models.BadRequestMessage
		_ = json.Unmarshal(responseData, &res)
		assert.Equal(suite.T(), http.StatusBadRequest, w.Code, test.name)
	}
	//
	//payload := &models.SignUpModel{
	//	Name:            "",
	//	Email:           "",
	//	Nickname:        "",
	//	ZipCode:         "",
	//	DateBirth:       time.Time{},
	//	Password:        "",
	//	ConfirmPassword: "",
	//}
	//
	//buf, _ := json.Marshal(payload)
	//
	//req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/auth/sign-up", bytes.NewReader(buf))
	//suite.router.ServeHTTP(w, req)
	//
	////mockResponse := "Its Alive and Kicking!"
	//responseData, _ := io.ReadAll(w.Body)
	//var res models.BadRequestMessage
	//err := json.Unmarshal(responseData, &res)
	//if err != nil {
	//	return
	//}
	//
	//assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *AuthControllerSuite) TestShouldCreateUserWithSuccess() {

	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)

	model := &models.SignUpModel{
		Name:            "Marcos",
		Email:           "marcos.mota287@gmail.com",
		Nickname:        "mota287",
		ZipCode:         "14085-058",
		Password:        "q1w2e3r4",
		ConfirmPassword: "q1w2e3r4",
	}

	buf, _ := json.Marshal(model)

	req, _ := http.NewRequestWithContext(ctx, "POST", "/api/v1/auth/sign-up", bytes.NewReader(buf))
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func TestAuthControllerSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerSuite))
}
