package handlers

import (
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleLanding(c echo.Context) error {
	page := templates.LandingPage()
	return render(200, c, templates.Layout(page))
}
