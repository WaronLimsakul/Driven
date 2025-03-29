package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleDoneTaskDay(c echo.Context) error {
	updatedTask, err := h.doneTaskForUser(c)
	if err != nil {
		return err
	}

	return render(http.StatusCreated, c, templates.DoneBigTaskResponse(updatedTask))
}
