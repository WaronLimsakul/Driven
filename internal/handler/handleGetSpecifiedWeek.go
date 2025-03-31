package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetSpecifiedWeek(c echo.Context) error {
	dateStr := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "cannot parse the requested date")
	}

	userID := c.Request().Header.Get("Driven-userID")
	if userID == "" {
		c.Logger().Error("at HandleGetSpecifiedWeek. NO user id header")
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Errorf("at HandleGetSpecifiedWeek: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	monday, sunday := tasks.GetWeekRange(date)

	getTasksParams := database.GetUserTasksWeekParams{
		OwnerID: userUUID,
		Date:    monday,
		Date_2:  sunday,
	}

	// They might not have task that week
	usersTasks, _ := h.Db.GetUserTasksWeek(c.Request().Context(), getTasksParams)
	groupedTasks := tasks.GroupTaskDate(usersTasks)
	return render(
		http.StatusOK, c,
		templates.Week(groupedTasks, monday))
}
