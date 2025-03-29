package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostTaskDay(c echo.Context) error {
	newTask, err := h.CreateTaskForUser(c)
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Reswap", "beforeend")
	c.Response().Header().Add("HX-Retarget", "#big-tasks-column")

	return render(http.StatusCreated, c, templates.BigTask(newTask))
}
