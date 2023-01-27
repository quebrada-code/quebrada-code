package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"quebrada_api/internal/app/models"
	"quebrada_api/internal/domain/services"
	"quebrada_api/internal/mapper"
	"quebrada_api/pkg/identity"
)

type AuthController struct {
	controllerBase
	authService  services.IAuthService
	tokenManager identity.TokenManager
}

func NewAuthController(
	authService services.IAuthService,
	tokenManager identity.TokenManager) *AuthController {
	return &AuthController{
		authService:  authService,
		tokenManager: tokenManager,
	}
}

// SignIn @Summary SignIn
// @Tags Auth
// @Description admin sign in
// @ModuleID SignIn
// @Accept  json
// @Produce  json
// @Param input body models.EmailCredential true "sign up info"
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /auth/sign-in [post]
func (a AuthController) SignIn(c *gin.Context) {

	model, err := ValidateModel[models.EmailCredential](c)
	if err != nil {
		errors, _ := json.Marshal(err)
		a.BadRequest(c, errors)
	}

	user, err := a.authService.SignInWithEmail(model.Email, model.Password)
	if err != nil {
		return
	}

	token, err := a.tokenManager.GenerateToken(user)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.TokenResponse{AccessToken: token})
}

// SignUp @Summary SignUp
// @Tags Auth
// @Description admin sign in
// @ModuleID SignUp
// @Accept  json
// @Produce  json
// @Param input body models.SignUpModel true "sign up info"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /auth/sign-up [post]
func (a AuthController) SignUp(c *gin.Context) {

	model, err := ValidateModel[models.SignUpModel](c)
	if err != nil {
		return
	}

	if model.Password != model.ConfirmPassword {
		c.AbortWithStatusJSON(
			400,
			models.BadRequestMessage{
				Errors: map[string]error{
					"password": errors.New("senhas não correspondem"),
				},
			},
		)
		return
	}

	entity := mapper.ToEntity(model)

	err = a.authService.CreateUser(entity)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Usuário salvo com sucesso!"})
}

// VerificationUser @Summary VerificationUser
// @Tags Auth
// @Description Verify user
// @ModuleID VerificationUser
// @Accept  json
// @Produce  json
// @Param input body models.UserVerificationCode true "Verification Code"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /auth/user-verification [post]
func (a AuthController) VerificationUser(c *gin.Context) {

	model, err := ValidateModel[models.UserVerificationCode](c)
	if err != nil {
		return
	}

	err = a.authService.ConfirmEmail(model.UserId, model.Code)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Usuário confirmado com sucesso!"})
}

// ResetPassword @Summary ResetPassword
// @Tags Auth
// @Description reset user password
// @ModuleID ResetPassword
// @Accept  json
// @Produce  json
// @Param input body models.ResetPassword true "sign up info"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /auth/reset-password [post]
func (a AuthController) ResetPassword(c *gin.Context) {

	model, err := ValidateModel[models.ResetPassword](c)
	if err != nil {
		return
	}

	token, err := a.authService.GeneratePasswordResetToken(model.Email)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, token)
}

// ConfirmResetPassword @Summary ConfirmResetPassword
// @Tags Auth
// @Description reset user password
// @ModuleID ConfirmResetPassword
// @Accept  json
// @Produce  json
// @Param input body models.ConfirmResetPassword true "sign up info"
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /auth/confirm-reset-password [post]
func (a AuthController) ConfirmResetPassword(c *gin.Context) {

	model, err := ValidateModel[models.ConfirmResetPassword](c)
	if err != nil {
		return
	}

	err = a.authService.ResetPassword(model.Email, model.Token, model.NewPassword)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Writer.WriteHeader(204)
}
