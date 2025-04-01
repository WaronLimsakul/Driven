package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleUndoneTaskWeek(c echo.Context) error {
	dayTasks, err := h.undoneTaskForUser(c)
	if err != nil {
		return err
	}

	return render(http.StatusCreated, c,
		templates.SmallTasksColumn(dayTasks, dayTasks[0].Date))
}
