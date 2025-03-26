package main

import (
	"log"

	"github.com/WaronLimsakul/Driven/internal/config"
	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/WaronLimsakul/Driven/internal/middlewares"
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

	serverMiddleware, err := middlewares.NewServerMiddlware()
	if err != nil {
		log.Fatal(err)
	}

	if serverConfig.Env == "development" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method} | uri=${uri} | status=${status} | err=${error}\n",
		}))
	}

	// give static files server (css, htmx script)
	e.Static("/static", "static")

	// server one file
	e.File("/favicon.ico", "static/ico/favicon.ico")

	e.GET("/", handlers.HandleLanding)
	e.GET("/home", serverMiddleware.AuthMiddleware(handlers.HandleGetHome))
	e.GET("/week", handlers.HandleGetWeek) // need middleware
	e.GET("/day", handlers.HandleGetDay)
	e.GET("/signin", handlers.HandleGetSignin)
	e.GET("/signup", handlers.HandleGetSignUp)

	e.POST("/signup", serverDBHandlers.HandlePostSignUp)
	e.POST("/signin", serverDBHandlers.HandlePostSignin)

	e.GET("/error", handlers.HandleLandError)
	e.GET("/*", handlers.HandleNotFound) // for not found page

	e.Start(serverConfig.Port)
}
