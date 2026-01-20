package http

import (
	"backend/src/internal/http/controllers"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New(controllers *controllers.Controller) error {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	// TODO not secure, just do it for my frontend
	e.Use(middleware.CORS())

	defineRoutes(e, controllers)

	err := e.Start(":8080")
	if err != nil {
		slog.Error("Error starting server")
		return err
	}

	return nil
}

func defineRoutes(e *echo.Echo, c *controllers.Controller) {
	project := e.Group("/project")
	project.GET("", c.Project.ListProjects)
	project.GET("/:id", c.Project.GetProject)
	project.POST("", c.Project.CreateProject)
	// project.PUT("/:id", c.Project.UpdateProject)
	project.DELETE("/:id", c.Project.DeleteProject)
}
