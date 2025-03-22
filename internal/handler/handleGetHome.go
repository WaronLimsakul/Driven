package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleGetHome(c echo.Context) error {
	page := templates.Home()
	return render(http.StatusOK, c, templates.AppLayout(page))
}
