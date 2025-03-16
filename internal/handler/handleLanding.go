package handler

import (
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleLanding(c echo.Context) error {
	return render(200, c, templates.Layout(templates.Home()))
}
