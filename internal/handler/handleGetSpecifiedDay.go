package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetSpecifiedDay(c echo.Context) error {
	inputDate := c.Param("date")
	date, err := time.Parse(time.DateOnly, inputDate)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't parse specified day")
	}

	userID := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Printf("At HandleGetToday: couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "somethign went wrong in the server")
	}

	getTasksParams := database.GetTaskByDateParams{
		OwnerID: userUUID,
		Date:    date,
	}

	todaysTasks, err := h.Db.GetTaskByDate(c.Request().Context(), getTasksParams)

	scrollTarget := c.Request().Header.Get("scrollTarget")
	if scrollTarget != "" {
		scrollHeaderVal := fmt.Sprintf("{\"scrollToTask\": \"%s\"}", scrollTarget)
		c.Response().Header().Set("HX-Trigger-After-Settle", scrollHeaderVal)
	}

	return render(http.StatusOK, c, templates.Day(todaysTasks, date))
}
