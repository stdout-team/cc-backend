package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) GetAllInterests(c *gin.Context) {
	resp, err := controller.evService.GetInterests()

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
