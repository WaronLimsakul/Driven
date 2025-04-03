package handlers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/WaronLimsakul/Driven/internal/database"
	_ "github.com/lib/pq" // when sql.Open(), it needs this driver
)

// handlers that need access to database
type DBHandler struct {
	Db        *database.Queries
	JWTSecret string
	Env       string
}

func NewDBHandler() (DBHandler, error) {

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return DBHandler{}, fmt.Errorf("database URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return DBHandler{}, err
	}
	queries := database.New(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return DBHandler{}, fmt.Errorf("couldn't find a Jwt secret config")
	}

	env := os.Getenv("ENV")
	if env == "" {
		return DBHandler{}, fmt.Errorf("couldn't get env config")
	}

	return DBHandler{
		Db:        queries,
		JWTSecret: jwtSecret,
		Env:       env,
	}, nil
}
