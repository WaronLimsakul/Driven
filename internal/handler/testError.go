package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func TestError(c echo.Context) error {
	page := templates.ServerErrorPage()
	completeComp := templates.Layout(page)
	return render(http.StatusInternalServerError, c, completeComp)
}
