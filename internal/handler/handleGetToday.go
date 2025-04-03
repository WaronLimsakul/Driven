package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetToday(c echo.Context) error {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	userID := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Printf("At HandleGetToday: couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "somethign went wrong in the server")
	}

	getTasksParams := database.GetTaskByDateParams{
		OwnerID: userUUID,
		Date:    today,
	}

	// don't care error, still want to render when today's task not found
	todaysTasks, _ := h.Db.GetTaskByDate(c.Request().Context(), getTasksParams)

	return render(http.StatusOK, c, templates.Day(todaysTasks, today))
}
