package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) CountMeIn(c *gin.Context) {
	resp, err := controller.evService.CountMeIn(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
