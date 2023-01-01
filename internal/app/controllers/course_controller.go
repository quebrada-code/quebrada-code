package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CourseController struct {
	controllerBase
}

func NewCourseController() *CourseController {
	return &CourseController{}
}

// GetCourses @Summary GetCourses
// @Tags Course
// @Description Get All Course
// @ModuleID GetCourses
// @Accept  json
// @Produce  json
// @Success 200 {object} models.CourseModel
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /course [get]
func (a CourseController) GetCourses(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}

// GetCoursesById @Summary GetCoursesById
// @Tags Course
// @Description Get course by id
// @ModuleID GetCoursesById
// @Accept  json
// @Produce  json
// @Success 200 {object} models.CourseModel
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /course [get]
func (a CourseController) GetCoursesById(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}

// CreateCourse @Summary CreateCourse
// @Tags Course
// @Description Create course
// @ModuleID CreateCourse
// @Accept  json
// @Produce  json
//
// @Success 200 {object} models.CourseModel
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /course [get]
func (a CourseController) CreateCourse(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}

// UpdateCourse @Summary UpdateCourse
// @Tags Course
// @Description Update course
// @ModuleID UpdateCourse
// @Accept  json
// @Produce  json
//
// @Success 200 {object} models.CourseModel
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /course [get]
func (a CourseController) UpdateCourse(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}

// DeleteCourse @Summary DeleteCourse
// @Tags Course
// @Description Delete course by Id
// @ModuleID DeleteCourse
// @Accept  json
// @Produce  json
//
// @Success 204 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /course [get]
func (a CourseController) DeleteCourse(c *gin.Context) {

	c.JSON(http.StatusOK, "")
}
