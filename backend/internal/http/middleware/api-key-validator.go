package middleware

import (
	"crypto/subtle"
	"log/slog"
	"net/http"
	"os"

	"backend/internal/http/messages"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Code         int    `json:"code"`
	ShortMessage string `json:"short_message"`
	Message      string `json:"message"`
}

func ApiKeyValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			slog.Error("API_KEY environment variable is not set")
			return c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:         http.StatusInternalServerError,
				ShortMessage: messages.SERVER_ERROR_SHORT,
				Message:      messages.SERVER_ERROR,
			})
		}

		providedKey := c.Request().Header.Get("X-API-KEY")
		if providedKey == "" {
			slog.Error("Request with missing API Key.")
			return c.JSON(http.StatusUnauthorized, ErrorResponse{
				Code:         http.StatusUnauthorized,
				ShortMessage: messages.NOT_AUTHORISED_SHORT,
				Message:      messages.NOT_AUTHORISED,
			})
		}

		// More secure, constant time of comparing
		if subtle.ConstantTimeCompare([]byte(providedKey), []byte(apiKey)) == 1 {
			return next(c)
		}

		slog.Error("Invalid API Key provided.")
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Code:         http.StatusUnauthorized,
			ShortMessage: messages.NOT_AUTHORISED_SHORT,
			Message:      messages.NOT_AUTHORISED,
		})
	}
}
