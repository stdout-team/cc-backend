package controllers

import (
	"cc/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) AddEvent(c *gin.Context) {
	input := &models.EventInput{}

	if err := c.BindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	event, err := controller.evService.AddEvent(input)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, event)
}
