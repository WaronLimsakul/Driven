package handlers

import (
	"fmt"
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostTaskWeek(c echo.Context) error {
	newTask, statusCode, err := h.CreateTaskForUser(c)
	if err != nil {
		return c.String(statusCode, err.Error())
	}
	// Retarget to add the end of task column
	c.Response().Header().Add("HX-Reswap", "beforeend")
	c.Response().Header().Add(
		"HX-Retarget",
		fmt.Sprintf("[id='%s']", templates.GetSmallTasksColumnID(newTask.Date)))

	return render(http.StatusCreated, c, templates.SmallTask(newTask))
}
