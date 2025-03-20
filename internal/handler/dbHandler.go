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
	Db *database.Queries
}

func NewDBHandler() (DBHandler, error) {

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return DBHandler{}, fmt.Errorf("Database URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return DBHandler{}, err
	}
	queries := database.New(db)

	return DBHandler{Db: queries}, nil
}
