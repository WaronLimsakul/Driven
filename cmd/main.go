package main

import (
	"fmt"
	"log"
	"os"

	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type serverConfig struct {
	port string
	env  string
}

func newServerConfig() (serverConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return serverConfig{}, fmt.Errorf("Couldn't load .env file")
	}
	port := os.Getenv("HOST")
	if port == "" {
		return serverConfig{}, fmt.Errorf("Couldn't get host number")
	}
	env := os.Getenv("ENV")
	if env == "" {
		return serverConfig{}, fmt.Errorf("Couldn't get env config")
	}

	return serverConfig{
		port: port,
		env:  env,
	}, nil
}

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
	e.File("/favicon.ico", "static/ico/favicon.ico")

	e.GET("/", handlers.HandleLanding)
	e.GET("/week", handlers.HandleGetWeek)
	e.GET("/day", handlers.HandleGetDay)
	e.GET("/signin", handlers.HandleGetSignin)
	e.GET("/signup", handlers.HandleGetSignUp)

	e.Start(config.port)
}
