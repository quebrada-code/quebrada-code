package controllers

import (
	"github.com/gin-gonic/gin"
	"quebrada_api/internal/app/models"
)

type controllerBase struct{}

func (a controllerBase) BadRequest(c *gin.Context, errors interface{}) {
	c.AbortWithStatusJSON(400, errors)
}

func ValidateModel[T models.IModel](c *gin.Context) (T, error) {

	var model T
	if err := c.BindJSON(&model); err != nil {
		return model, err
	}

	err := model.Validate()
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
	}

	return model, nil

}
