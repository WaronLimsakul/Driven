package handlers

import (
	"net/http"

	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleDoneTaskWeek(c echo.Context) error {
	updatedTask, err := h.doneTaskForUser(c)

	if err != nil {
		return err
	}

	taskWeekDay := tasks.GetWeekDayStr(updatedTask.Date)

	return render(
		http.StatusCreated, c,
		templates.DoneSmallTaskResponse(updatedTask, taskWeekDay))
}
