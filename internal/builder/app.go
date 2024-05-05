package builder

import (
	"go-echo/internal/http/handler"
	"go-echo/internal/http/router"
	"go-echo/internal/repository"
	"go-echo/internal/service"
	"go-echo/pkg/route"
	"go-echo/pkg/token"

	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPublicRoutes(userHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB) []*route.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, nil)
	userHandler := handler.NewUserHandler(userService)
	return router.AppPrivateRoutes(userHandler)
}
