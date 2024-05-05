package router

import (
	"go-echo/pkg/route"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthPublicRoutes() []*route.Route {
	return []*route.Route{
		{
			Method: http.MethodGet,
			Path:   "/login",
			Handler: func(c echo.Context) error {
				return c.JSON(http.StatusOK, "Hello Haqiqi Ganteng")
			},
		},
	}
}

func AuthPrivateRoutes() []*route.Route {
	return nil
}
