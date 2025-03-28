package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/database"
	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleDoneTask(c echo.Context) error {
	taskIDStr := c.Param("id")
	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.Logger().Printf("couldn't parse task id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	userIDStr := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.Logger().Printf("couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	doneTaskArgs := database.DoneTaskByIDParams{
		ID:      taskUUID,
		OwnerID: userUUID,
	}

	// check if user is correct in one query
	updatedTask, err := h.Db.DoneTaskByID(c.Request().Context(), doneTaskArgs)
	if err != nil {
		c.Logger().Printf("couldn't update task id: %v", err)
		return c.String(http.StatusUnauthorized, "something went wrong")
	}

	taskWeekDay := tasks.GetWeekDayStr(updatedTask.Date)

	return render(
		http.StatusCreated, c,
		templates.DoneTaskResponse(updatedTask, taskWeekDay))
}
