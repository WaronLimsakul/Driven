package handlers

import (
	"net/http"

	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleUndoneTaskWeek(c echo.Context) error {
	dayTasks, err := h.undoneTaskForUser(c)
	if err != nil {
		return err
	}

	weekDay := tasks.GetWeekDayStr(dayTasks[0].Date)
	return render(http.StatusCreated, c, templates.SmallTasksColumn(dayTasks, weekDay))
}
