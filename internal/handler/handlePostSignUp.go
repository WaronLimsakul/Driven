package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostSignUp(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := h.Db.GetUserByEmail(c.Request().Context(), email)

	// there is existed email
	if err == nil {
		return c.String(http.StatusConflict, "*Email already existed")
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot hash password")
	}

	params := database.CreateUserParams{
		ID:             uuid.New(),
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	_, err = h.Db.CreateUser(c.Request().Context(), params)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot create new user")
	}

	return render(http.StatusCreated, c, templates.SignUpSuccessMessage())
}
