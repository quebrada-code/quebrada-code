package router

import "github.com/gin-gonic/gin"

func (h *Router) initCourseRouters(api *gin.RouterGroup) {
	users := api.Group("/course")
	{
		users.POST("/sign-in", h.AuthController.SignIn)
		users.POST("/sign-up", h.AuthController.SignUp)
		users.POST("/user-verification", h.AuthController.VerificationUser)
		users.POST("/reset-password", h.AuthController.ResetPassword)
		users.POST("/confirm-reset-password", h.AuthController.ConfirmResetPassword)
	}
}
