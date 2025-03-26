package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleGetHome(c echo.Context) error {
	page := templates.Home()
	script := templates.HomePageScript()
	layout := templates.InjectedAppLayout
	return render(http.StatusOK, c, layout(page, script))
}
