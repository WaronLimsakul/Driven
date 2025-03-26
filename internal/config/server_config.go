package config

import (
	"fmt"
	"os"
)

type ServerConfig struct {
	Port string
	Env  string
}

func NewServerConfig() (ServerConfig, error) {
	port := os.Getenv("HOST")
	if port == "" {
		return ServerConfig{}, fmt.Errorf("Couldn't get host number")
	}
	env := os.Getenv("ENV")
	if env == "" {
		return ServerConfig{}, fmt.Errorf("Couldn't get env config")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return ServerConfig{}, fmt.Errorf("Couldn't get database url")
	}

	return ServerConfig{
		Port: port,
		Env:  env,
	}, nil
}
