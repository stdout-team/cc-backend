package controllers

import (
	"cc/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *Controller) EventsLookup(c *gin.Context) {
	apiReq := &models.EventsLookupAPIRequest{}

	if err := c.BindQuery(apiReq); err != nil {
		c.Error(err)
		return
	}

	req, err := models.NewEventsLookupRequest(apiReq)
	if err != nil {
		c.Error(err)
		return
	}

	resp, err := controller.evService.Lookup(req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
