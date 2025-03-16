package main

import (
	"github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	// give static files server (css, htmx script)
	e.Static("/static", "static")

	e.GET("/", handler.HandleLanding)

	e.Start(":8080")
}
