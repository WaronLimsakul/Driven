package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleGetSignin(c echo.Context) error {
	page := templates.SigninPage()
	script := templates.SignInScript()
	complete := templates.InjectedLayout(page, script)
	return render(http.StatusOK, c, complete)
}
