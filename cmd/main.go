package main

import (
	"log"

	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config, err := newServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	if config.env == "development" {
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

	e.POST("/signup", handlers.HandlePostSignUp)

	e.Start(config.port)
}
