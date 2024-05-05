package builder

import (
	"go-echo/internal/http/router"
	"go-echo/pkg/route"
)

func BuildAuthPublicRoutes() []*route.Route {
	return router.AuthPublicRoutes()
}

func BuildAuthPrivateRoutes() []*route.Route {
	return router.AuthPrivateRoutes()
}
