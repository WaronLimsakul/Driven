package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func HandleNotFound(c echo.Context) error {
	page := templates.NotFoundPage()
	completeComp := templates.Layout(page)
	return render(http.StatusNotFound, c, completeComp)
}
