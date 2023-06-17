package router

import (
	"github.com/gin-gonic/gin"
	"quebrada_api/internal/app/controllers"
)

// @title Quebrada Code API
// @version 1.0
// @description REST API for Quebrada Code API

// @host localhost:9090
// @BasePath /api/v1/

type Router struct {
	AuthController   controllers.AuthController
	CourseController controllers.CourseController
}

func (h *Router) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
	}
}
