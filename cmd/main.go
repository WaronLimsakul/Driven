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
	e.GET("/week", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleGetWeek))
	e.GET("/week/:date", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleGetSpecifiedWeek))
	e.GET("/day", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleGetToday))
	e.GET("/day/:date", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleGetSpecifiedDay))
	e.GET("/signin", handlers.HandleGetSignin)
	e.GET("/signup", handlers.HandleGetSignUp)

	e.POST("/signup", middlewares.HXFilterMiddleWare(serverDBHandlers.HandlePostSignUp))
	e.POST("/signin", middlewares.HXFilterMiddleWare(serverDBHandlers.HandlePostSignin))
	e.POST("/signout", serverDBHandlers.HandlePostSignOut)

	e.POST("/tasks/week", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandlePostTaskWeek))
	e.PUT("/tasks/week/:id/done", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleDoneTaskWeek))
	e.PUT("/tasks/week/:id/undone", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleUndoneTaskWeek))

	e.POST("/tasks/day", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandlePostTaskDay))
	e.PUT("/tasks/day/:id/done", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleDoneTaskDay))
	e.PUT("/tasks/day/:id/undone", serverMiddleware.HXAuthMiddleware(serverDBHandlers.HandleUndoneTaskDay))
	e.PUT("/tasks/day/:id/keys", serverMiddleware.AuthMiddleware(serverDBHandlers.HandlePutTaskKeys))

	e.DELETE("/tasks/:id", serverMiddleware.AuthMiddleware(serverDBHandlers.HandleDeleteTask))

	e.GET("/error", handlers.HandleLandError)
	e.GET("/*", handlers.HandleNotFound) // for not found page

	err = e.Start(serverConfig.Port)
	if err != nil {
		log.Fatalf("Couldn't start the server: %v", err)
	}
}
