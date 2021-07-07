package routes

import (
	"net/http"

	"github.com/Aldiwildan77/rust-notifier-api/api"
	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/routes"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Iam fine thanks")
	})

	routes.NewRoute(api.ProvideHandler()).Init(e)
}