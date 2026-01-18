package api

import (
	"backend/db/minitec_db"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(queries *minitec_db.Queries) error {
	e := echo.New()

	// TODO not secure, just do it for my frontend
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	err := e.Start(":8080")
	if err != nil {
		slog.Error("Error starting server")
		return err
	}

	return nil
}
