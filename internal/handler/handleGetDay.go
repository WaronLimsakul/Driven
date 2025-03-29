package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetDay(c echo.Context) error {

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	userID := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Printf("At HandleGetDay: couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "somethign went wrong in the server")
	}

	getTasksParams := database.GetTaskByDateParams{
		OwnerID: userUUID,
		Date:    today,
	}

	todaysTasks, err := h.Db.GetTaskByDate(c.Request().Context(), getTasksParams)

	return render(http.StatusOK, c, templates.Day(todaysTasks, today))
}
