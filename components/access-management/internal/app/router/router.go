package router

import (
	"github.com/gin-gonic/gin"
	"quebrada_api/internal/app/controllers"
)

// @title Authentication API
// @version 1.0
// @description REST API for Quebrada Code API

// @host localhost:9090
// @BasePath /api/v1/

type Router struct {
	AuthController    controllers.AuthController
	ProblemController controllers.ProblemController
}

func (h *Router) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initAuthRoutes(v1)
		h.initProblemRoutes(v1)
	}
}
