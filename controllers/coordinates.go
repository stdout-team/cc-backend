package controllers

import (
	"cc/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResolveCoordinates(c *gin.Context) {
	type validator struct {
		Address string `binding:"required" form:"address"`
	}

	req := &validator{}

	if err := c.BindQuery(req); err != nil {
		c.Error(err)
		return
	}

	coordinates, err := services.ResolveCoordinatesByStreetAddress(req.Address)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"lat": coordinates[0], "lon": coordinates[1]})
}
