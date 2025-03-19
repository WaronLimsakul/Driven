package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleGetSignin(c echo.Context) error {
	return render(http.StatusOK, c, templates.Layout(templates.SigninPage()))
}
