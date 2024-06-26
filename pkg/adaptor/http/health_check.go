package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthCheckResponse struct {
	Message string `json:"message"`
}

func healthCheck(c echo.Context) error {
	message := fmt.Sprintf("Hello! you've requested: %s", c.Path())
	return c.JSON(
		http.StatusOK,
		healthCheckResponse{
			Message: message,
		},
	)
}
