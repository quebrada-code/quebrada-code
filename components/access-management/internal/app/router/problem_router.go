package router

import "github.com/gin-gonic/gin"

func (h *Router) initProblemRoutes(api *gin.RouterGroup) {
	problem := api.Group("/problems")
	{
		problem.GET("/:pk", h.ProblemController.GetById)
		problem.GET("/", h.ProblemController.GetAll)
		problem.POST("/", h.ProblemController.Create)
		problem.POST("/submit", h.ProblemController.SubmitSolution)

	}
}
