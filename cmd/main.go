package main

import (
	"log"

	"github.com/WaronLimsakul/Driven/internal/config"
	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/WaronLimsakul/Driven/internal/middlewares"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Couldn't load .env file: %v", err)
	}

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
	e.GET("/home", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleGetHome))
	e.GET("/week", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleGetWeek))
	e.GET("/day", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleGetDay))
	e.GET("/signin", handlers.HandleGetSignin)
	e.GET("/signup", handlers.HandleGetSignUp)

	e.POST("/signup", serverDBHandlers.HandlePostSignUp)
	e.POST("/signin", serverDBHandlers.HandlePostSignin)

	e.POST("/tasks", serverMiddleware.AuthMiddleware(serverDBHandlers.HandlePostTask))
	e.PUT("/tasks/week/:id/done", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleDoneTaskWeek))
	e.PUT("/tasks/week/:id/undone", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleUndoneTaskWeek))
	e.GET("/error", handlers.HandleLandError)
	e.GET("/*", handlers.HandleNotFound) // for not found page

	e.Start(serverConfig.Port)
}
