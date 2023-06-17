package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quebrada_api/internal/app/models"
	"quebrada_api/internal/domain/services"
	"quebrada_api/internal/mapper"
	"strconv"
)

type ProblemController struct {
	controllerBase
	problemService services.IProblemService
}

func NewProblemController(
	problemService services.IProblemService) *ProblemController {
	return &ProblemController{
		problemService: problemService,
	}
}

// GetAll  @Summary SignIn
// @Tags Problems
// @Description Get all problems
// @ModuleID GetAll
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /problems/ [get]
func (a ProblemController) GetAll(c *gin.Context) {

	problems, err := a.problemService.GetProblems()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, problems)
}

// GetById  @Summary SignIn
// @Tags Problems
// @Description Get problem by Id
// @ModuleID GetById
// @Param   pk     path    string     true        "Feature Dataset Primary Key"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /problems/{pk} [get]
func (a ProblemController) GetById(c *gin.Context) {

	idParam := c.Param("pk")
	problemId, err := strconv.ParseUint(idParam, 10, 0)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]string{"id": "Id não informado"})
		return
	}
	problem, err := a.problemService.GetProblemById(uint(problemId))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, models.ResponseModel{Data: problem})
}

// Create  @Summary SignIn
// @Tags Problems
// @Description Get problem by Id
// @ModuleID Create
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /problems/ [post]
func (a ProblemController) Create(c *gin.Context) {

	model, err := ValidateModel[models.CreateProblem](c)
	if err != nil {
		return
	}
	entity := mapper.ToProblem(model)
	err = a.problemService.CreateProblem(entity)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Problema salvo com sucesso!"})
}

// SubmitSolution  @Summary Submit solution for problem
// @Tags Problems
// @Description Submit solution for problem
// @ModuleID SubmitSolution
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /problems/submit [post]
func (a ProblemController) SubmitSolution(c *gin.Context) {

	model, err := ValidateModel[models.SubimtProblemModel](c)
	if err != nil {
		return
	}

	err = a.problemService.SubmitProblem(model.UserId, model.ProblemId, model.SolutionCode)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, models.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Message: "Problema salvo com sucesso!"})
}
