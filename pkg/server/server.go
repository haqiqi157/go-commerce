package server

import (
	"context"
	"fmt"
	"go-echo/pkg/response"
	"go-echo/pkg/route"
	"go-echo/pkg/token"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(serverName string, publicRoutes, privateRoutes []*route.Route, secretKey string) *Server {
	e := echo.New()
	//v1 := e.Group("api/v1")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Haqiqi Ganteng", nil))
	})

	v1 := e.Group(fmt.Sprintf("%s/api/v1", serverName))

	if len(publicRoutes) > 0 {
		for _, v := range publicRoutes {
			v1.Add(v.Method, v.Path, v.Handler)
		}
	}

	if len(privateRoutes) > 0 {
		for _, v := range privateRoutes {
			v1.Add(v.Method, v.Path, v.Handler, JWTProtection(secretKey))
		}
	}

	return &Server{e}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s)
}

func runServer(srv *Server) {
	go func() {
		err := srv.Start(":8080")
		log.Fatal(err)
	}()
}

func gracefulShutdown(srv *Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal("Server shutdown:", err)
		}
	}()
}

func JWTProtection(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "anda harus login untuk mengakses resource ini"))
		},
	})
}
