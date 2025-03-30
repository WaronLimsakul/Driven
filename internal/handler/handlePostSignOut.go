package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// 1. get refresh token and delete from db
// 2. delete every driven-related cookie from request
// 3. redirect back to sign in page
func (h DBHandler) HandlePostSignOut(c echo.Context) error {
	c.Response().Header().Set("HX-Redirect", "/signin")

	refreshTokenCookie, err := c.Cookie("driven-refresh-token")
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't find user refresh token")
	}

	refreshTokenCookie.Expires = time.Now()
	c.SetCookie(refreshTokenCookie)

	refreshTokenValue := refreshTokenCookie.Value
	err = h.Db.DeleteToken(c.Request().Context(), refreshTokenValue)
	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't delete token")
	}

	accessTokenCookie, err := c.Cookie("driven-jwt")
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't find access token")
	}

	accessTokenCookie.Expires = time.Now()
	c.SetCookie(accessTokenCookie)

	return c.String(http.StatusCreated, "user signed out")
}
