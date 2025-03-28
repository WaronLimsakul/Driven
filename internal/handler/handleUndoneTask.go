package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/database"
	tasks "github.com/WaronLimsakul/Driven/internal/task"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleUndoneTaskWeek(c echo.Context) error {
	userIDStr := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.Logger().Printf("at HandleUndoneTask: couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	taskIDStr := c.Param("id")
	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.Logger().Printf("at HandleUndoneTask: couldn't parse task id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	undoneTaskParams := database.UndoneTaskByIDParams{
		ID:      taskUUID,
		OwnerID: userUUID,
	}

	updatedTask, err := h.Db.UndoneTaskByID(c.Request().Context(), undoneTaskParams)
	if err != nil {
		return c.String(http.StatusUnauthorized, "couldn't undone task: user not the owner")
	}

	// Have to render entire column again because we don't know where to put it back
	taskDate := updatedTask.Date
	getDayTasksParam := database.GetTaskByDateParams{
		OwnerID: userUUID,
		Date:    taskDate,
	}
	dayTasks, err := h.Db.GetTaskByDate(c.Request().Context(), getDayTasksParam)

	if err != nil {
		c.Logger().Printf("couldn't get user tasks by date: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	weekDay := tasks.GetWeekDayStr(taskDate)
	return render(http.StatusCreated, c, templates.SmallTasksColumn(dayTasks, weekDay))
}
