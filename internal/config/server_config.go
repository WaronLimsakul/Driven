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
	port := os.Getenv("PORT")
	if port == "" {
		return ServerConfig{}, fmt.Errorf("couldn't get host number")
	}
	env := os.Getenv("ENV")
	if env == "" {
		return ServerConfig{}, fmt.Errorf("couldn't get env config")
	}

	return ServerConfig{
		Port: fmt.Sprintf(":%s", port),
		Env:  env,
	}, nil
}
