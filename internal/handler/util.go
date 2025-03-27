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

func AssignAuthCookies(c echo.Context, production bool, accessToken, refreshToken string) {
	AssignAccessTokenCookie(c, accessToken, production)
	AssignRefreshTokenCookie(c, refreshToken, production)
	return
}

func AssignAccessTokenCookie(c echo.Context, accessToken string, production bool) {
	// normal browser will bind this cookie to only the server it got cookie from
	accessTokenCookie := new(http.Cookie)
	accessTokenCookie.Name = "driven-jwt"
	accessTokenCookie.Value = accessToken
	accessTokenCookie.Expires = time.Now().Add(jwtExpireTime)
	accessTokenCookie.HttpOnly = true
	if production {
		accessTokenCookie.Secure = true
	}
	accessTokenCookie.Path = "/"
	c.SetCookie(accessTokenCookie)
	return
}

func AssignRefreshTokenCookie(c echo.Context, refreshToken string, production bool) {
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

func getWeekRange(today time.Time) (monday, sunday time.Time) {
	goBack := int(today.Weekday()) - 1 // Sunday is 0
	if goBack < 0 {
		goBack = 6
	}

	monday = today.AddDate(0, 0, -goBack)
	sunday = today.AddDate(0, 0, 6-goBack)
}
