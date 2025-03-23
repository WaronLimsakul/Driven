package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const jwtExpireTime time.Duration = 15 * time.Minute
const refreshExpireTime time.Duration = 7 * 24 * time.Hour

func render(status int, context echo.Context, templComp templ.Component) error {
	context.Response().Status = status
	return templComp.Render(context.Request().Context(), context.Response())
}

func createDoubleTokens(c echo.Context, userID uuid.UUID, secret string) (accessToken, refreshToken string, err error) {
	accessToken, err = auth.CreateJWT(userID, jwtExpireTime, secret)
	if err != nil {
		c.Logger().Errorf("Couldn't create access token: %v", err)
		return "", "", c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	refreshToken, err = auth.CreateRefreshToken()
	if err != nil {
		c.Logger().Errorf("Couldn't create refresh token: %v", err)
		return "", "", c.String(http.StatusInternalServerError, "Couldn't create refresh token")
	}

	return accessToken, refreshToken, nil
}

func assignAuthCookies(c echo.Context, production bool, accessToken, refreshToken string) {
	// normal browser will bind this cookie to only the server it got cookie from
	accessTokenCookie := new(http.Cookie)
	accessTokenCookie.Name = "driven-jwt"
	accessTokenCookie.Value = accessToken
	accessTokenCookie.Expires = time.Now().Add(jwtExpireTime)
	accessTokenCookie.HttpOnly = true
	if production {
		accessTokenCookie.Secure = true
	}
	c.SetCookie(accessTokenCookie)

	refreshTokenCookie := new(http.Cookie)
	refreshTokenCookie.Name = "driven-refresh-token"
	refreshTokenCookie.Value = refreshToken
	refreshTokenCookie.Expires = time.Now().Add(refreshExpireTime)
	refreshTokenCookie.HttpOnly = true
	if production {
		refreshTokenCookie.Secure = true
	}
	c.SetCookie(refreshTokenCookie)
	return
}
