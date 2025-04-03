package middlewares

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	handlers "github.com/WaronLimsakul/Driven/internal/handler"
	"github.com/labstack/echo/v4"
)

// this one has access to DB queries and JWT secret
type ServerMiddleware struct {
	Db        *database.Queries
	JWTSecret string
	Env       string
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

	env := os.Getenv("ENV")
	if env == "" {
		return ServerMiddleware{}, fmt.Errorf("Couldn't find ENV config")
	}

	return ServerMiddleware{
		Db:        queries,
		JWTSecret: jwtSecret,
		Env:       env,
	}, nil
}

// redirect to signin, if doesn't found refresh token
// refresh jwt if refresh token found but expired
// pass if jwt valid
func (m ServerMiddleware) AuthMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		refreshCookie, err := c.Cookie("driven-refresh-token")
		// it err != nil when don't found the cookie
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/signin")
		}

		refreshTokenValue := refreshCookie.Value
		refreshToken, err := m.Db.GetRefreshToken(c.Request().Context(), refreshTokenValue)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/signin")
		}

		err = auth.ValidateRefreshToken(refreshToken)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/signin")
		}

		jwtCookie, err := c.Cookie("driven-jwt")
		// there is refresh token but not jwt
		if err != nil {
			err = m.refreshJWT(refreshToken.Token, c)
			if err != nil {
				c.Logger().Printf("catch at 1: %v", err)
				return err
			}
			// fmt.Printf("token refreshed\n")
		}

		jwtCookie, err = c.Cookie("driven-jwt") // assign again after refresh
		if err != nil {
			c.Logger().Printf("at auth middleware, catch 2: %v", err)
			return err
		}

		userID, err, isExpired := auth.ValidateJWT(jwtCookie.Value, m.JWTSecret)
		if err != nil {
			c.Logger().Printf("At auth middleware (validate jwt): %v", err)
			return c.Redirect(http.StatusSeeOther, "/error")
		} else if isExpired {
			err = m.refreshJWT(refreshToken.Token, c)
			if err != nil {
				return err
			}
		}

		// set to have user id
		c.Request().Header.Set("Driven-userID", userID.String())
		return handler(c)
	}
}

// 1. find user id
// 2. pass it to our util function to get accesstoken
// 3. assign it in cookie
func (m ServerMiddleware) refreshJWT(refreshToken string, c echo.Context) error {
	userRefreshToken, err := m.Db.GetRefreshToken(c.Request().Context(), refreshToken)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/signin")
	}

	newAccessToken, err := middlewareCreateAccessToken(c, userRefreshToken.UserID, m.JWTSecret)
	if err != nil {
		return err // already set respond in helper
	}

	handlers.AssignAccessTokenCookie(c, newAccessToken, m.Env == "production")
	return nil
}

// combined version of HXFilter and Auth middleware
func (m ServerMiddleware) HXAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return HXFilterMiddleWare(m.AuthMiddleware(next))
}
