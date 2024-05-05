package main

import (
	"go-echo/configs"
	"go-echo/internal/builder"
	"go-echo/pkg/postgres"
	"go-echo/pkg/server"
	"go-echo/pkg/token"
)

func main() {
	//publicRoutes := builder.BuildAuthPublicRoutes()
	//privateRoutes := builder.BuildAuthPrivateRoutes()

	cfg, err := configs.NewConfig(".env")
	checkError(err)

	db, err := postgres.InitPostgres(&cfg.Postgres)
	checkError(err)

	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUseCase)
	privateRoutes := builder.BuildAppPrivateRoutes(db)

	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
