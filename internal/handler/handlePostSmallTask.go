package handlers

import (
	"fmt"
	"net/http"

	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostTaskWeek(c echo.Context) error {
	newTask, statusCode, err := h.CreateTaskForUser(c)
	if err != nil {
		return c.String(statusCode, err.Error())
	}
	weekDayStr := tasks.GetWeekDayStr(newTask.Date)
	c.Response().Header().Add("HX-Reswap", "beforeend")
	c.Response().Header().Add("HX-Retarget", fmt.Sprintf("#%s", weekDayStr))

	return render(http.StatusCreated, c, templates.SmallTask(newTask))
}
