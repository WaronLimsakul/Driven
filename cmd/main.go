package main

import (
	"log"

	"github.com/WaronLimsakul/Driven/internal/config"
	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	serverConfig, err := config.NewServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	serverDBHandlers, err := handlers.NewDBHandler()
	if err != nil {
		log.Fatal(err)
	}

	if serverConfig.Env == "development" {
		e.Use(middleware.Logger())
	}

	// give static files server (css, htmx script)
	e.Static("/static", "static")

	// server one file
	e.File("/favicon.ico", "static/ico/favicon.ico")

	e.GET("/", handlers.HandleLanding)
	e.GET("/week", handlers.HandleGetWeek)
	e.GET("/day", handlers.HandleGetDay)
	e.GET("/signin", handlers.HandleGetSignin)
	e.GET("/signup", handlers.HandleGetSignUp)

	e.POST("/signup", serverDBHandlers.HandlePostSignUp)
	e.POST("/signin", serverDBHandlers.HandlePostSignin)

	e.Start(serverConfig.Port)
}
