package controllers

import (
	"backend/internal/http/controllers/project"
	"backend/internal/http/services"
)

type Controller struct {
	Services *services.Services
	Project  *project.Project
}

func New(services *services.Services) *Controller {
	return &Controller{
		Project: project.New(services.Project),
	}
}
