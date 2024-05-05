package main

import (
	"go-echo/internal/builder"
	"go-echo/pkg/server"
)

func main() {
	publicRoutes := builder.BuildAuthPublicRoutes()
	privateRoutes := builder.BuildAuthPrivateRoutes()

	srv := server.NewServer("auth", publicRoutes, privateRoutes, "")
	srv.Run()
}
