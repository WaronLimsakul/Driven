package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetWeek(c echo.Context) error {
	monday, sunday := getWeekRange(time.Now().UTC())
	userID := c.Request().Header.Get("Driven-userID")
	if userID == "" {
		c.Logger().Error("at HandleGetWeek. NO user id header")
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Errorf("at HandleGetWeek: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	params := database.GetUserTasksWeekParams{
		OwnerID: userUUID,
		Date:    monday,
		Date_2:  sunday,
	}

	// assume that error just mean it doesn't found any tasks
	tasks, _ := h.Db.GetUserTasksWeek(c.Request().Context(), params)
	groupedTasks := groupTaskDate(tasks)

	return render(http.StatusOK, c, templates.Week(groupedTasks))
}
