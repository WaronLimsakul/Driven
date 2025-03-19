package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/joho/godotenv"
)

type serverConfig struct {
	port string
	env  string
	db   *database.Queries
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

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return serverConfig{}, fmt.Errorf("Couldn't get database url")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return serverConfig{}, fmt.Errorf("Couldn't open a database")
	}

	dbQueries := database.New(db)

	return serverConfig{
		port: port,
		env:  env,
		db:   dbQueries,
	}, nil
}
