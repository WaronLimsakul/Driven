package middlewares

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/labstack/echo/v4"
)

type ServerMiddleware struct {
	Db        *database.Queries
	JWTSecret string
}

func NewServerMiddlware() (ServerMiddleware, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return ServerMiddleware{}, fmt.Errorf("Database URL not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return ServerMiddleware{}, err
	}
	queries := database.New(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return ServerMiddleware{}, fmt.Errorf("Couldn't find a Jwt secret config")
	}

	return ServerMiddleware{
		Db:        queries,
		JWTSecret: jwtSecret,
	}, nil
}

func (m ServerMiddleware) AuthMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		refreshCookie, err := c.Cookie("driven-refresh-token")
		// it err != nil when don't found the cookie
		if err != nil {
			c.Request().Header.Add("HX-Redirect", "/signin")
			return c.String(http.StatusUnauthorized, "Refresh token not found in cookie")
		}

		refreshTokenValue := refreshCookie.Value
		refreshToken, err := m.Db.GetRefreshToken(c.Request().Context(), refreshTokenValue)
		if err != nil {
			c.Request().Header.Add("HX-Redirect", "/signin")
			return c.String(http.StatusUnauthorized, "Refresh token not found in DB")
		}

		err = auth.ValidateRefreshToken(refreshToken)
		if err != nil {
			c.Logger().Errorf("At auth middleware: %v", err)
			c.Request().Header.Add("HX-Redirect", "/signin")
			return c.String(http.StatusUnauthorized, "Refresh token invalid")
		}

		jwtCookie, err := c.Cookie("driven-jwt")
		// there is refresh token but not jwt
		if err != nil {
			// write a function that create a new jwt and assign it for req.
		}

		return handler(c)
	}
}
