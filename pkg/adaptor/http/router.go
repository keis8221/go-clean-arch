package http

import (
	"github.com/keis8221/go-clean-arch/pkg/infra/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	healthCheckGroup := e.Group(healthCheckRoot)
	{
		reactivePath := ""
		healthCheckGroup.GET(reactivePath, healthCheck)
	}

	database.NewPsqlConnector()

	return e
}
