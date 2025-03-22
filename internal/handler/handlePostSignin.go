package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/auth"
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

	err = assignAuthCookies(c, user.ID, h.JWTSecret, h.Env == "production")
	if err != nil {
		return err
	}

	return render(http.StatusCreated, c, templates.SignInSuccessMessage())
}
