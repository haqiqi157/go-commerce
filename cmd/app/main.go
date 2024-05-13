package main

import (
	"go-echo/configs"
	"go-echo/internal/builder"
	"go-echo/pkg/cache"
	"go-echo/pkg/encrypt"
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

	redisDB := cache.InitCache(&cfg.Redis)

	tokenUseCase := token.NewTokenUseCase(cfg.JWT.SecretKey)
	encryptTool := encrypt.NewEncryptTool(cfg.Encrypt.SecretKey, cfg.Encrypt.IV)

	publicRoutes := builder.BuildAppPublicRoutes(db, tokenUseCase, encryptTool)
	privateRoutes := builder.BuildAppPrivateRoutes(db, redisDB, encryptTool)

	srv := server.NewServer("app", publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
