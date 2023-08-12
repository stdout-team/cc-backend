package controllers

import (
	"cc/services"
)

type Controller struct {
	evService *services.EventsService
}

func NewController(evService *services.EventsService) *Controller {
	return &Controller{evService: evService}
}
