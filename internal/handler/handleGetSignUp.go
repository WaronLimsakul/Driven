package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleGetSignUp(c echo.Context) error {
	signUpPageScript := templates.SignUpScript()
	page := templates.SignUpPage()
	component := templates.InjectedLayout(page, signUpPageScript)
	return render(http.StatusOK, c, component)
}
