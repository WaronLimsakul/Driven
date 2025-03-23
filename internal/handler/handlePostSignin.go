package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h *DBHandler) HandlePostSignin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "" || password == "" {
		return c.String(http.StatusUnauthorized, "Invalid user or password")
	}

	user, err := h.Db.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid user or password")
	}

	err = auth.ValidatePassword(password, user.HashedPassword)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Invalid user or password")
	}

	accessToken, refreshToken, err := createDoubleTokens(c, user.ID, h.JWTSecret)
	if err != nil {
		return err
	}

	assignAuthCookies(c, h.Env == "production", accessToken, refreshToken)
	refreshTokenParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiredAt: time.Now().Add(refreshExpireTime),
	}

	_, err = h.Db.CreateRefreshToken(c.Request().Context(), refreshTokenParams)
	if err != nil {
		c.Logger().Errorf("At handlePostSignin, cannot add refresh token to db: %v", err)
		return c.String(500, "Something wen wrong")
	}

	return render(http.StatusCreated, c, templates.SignInSuccessMessage())
}
