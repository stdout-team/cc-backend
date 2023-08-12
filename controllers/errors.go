package controllers

import (
	"cc/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ErrorsMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	resp := &models.ErrorResponse{Timestamp: time.Now(), Errors: []string{}}

	for _, err := range c.Errors {
		resp.Errors = append(resp.Errors, err.Error())
	}

	c.JSON(http.StatusBadRequest, resp)
}
