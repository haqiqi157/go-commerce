package router

import (
	"go-echo/internal/http/handler"
	"go-echo/pkg/route"
	"net/http"
)

func AppPublicRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
	}
}

func AppPrivateRoutes(userHandler handler.UserHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: userHandler.FindAllUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  http.MethodPut,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: userHandler.FindUserByID,
		},
	}
}
