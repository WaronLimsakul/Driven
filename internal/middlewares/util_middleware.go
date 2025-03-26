package middlewares

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const jwtExpireTime time.Duration = 15 * time.Minute
const refreshExpireTime time.Duration = 7 * 24 * time.Hour

func middlewareCreateAccessToken(c echo.Context, userID uuid.UUID, secret string) (string, error) {
	accessToken, err := auth.CreateJWT(userID, jwtExpireTime, secret)
	if err != nil {
		c.Logger().Errorf("Couldn't create access token: %v", err)
		return "", c.Redirect(http.StatusSeeOther, "/error")
	}
	return accessToken, nil
}

// don't know if we need this
func middlwareCreateRefreshToken(c echo.Context) (string, error) {
	refreshToken, err := auth.CreateRefreshToken()
	if err != nil {
		c.Logger().Errorf("Couldn't create refresh token: %v", err)
		return "", c.Redirect(http.StatusSeeOther, "/error")
	}
	return refreshToken, err
}
