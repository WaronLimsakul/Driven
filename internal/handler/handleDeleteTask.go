package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *DBHandler) HandleDeleteTask(c echo.Context) error {
	taskID := c.Param("id")
	if taskID == "" {
		return c.String(http.StatusBadRequest, "no task id")
	}

	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return c.String(http.StatusBadRequest, "cannot parse task id")
	}

	userID := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userID)

	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't parse user id")
	}

	params := database.DeleteTaskByIDParams{
		ID:      taskUUID,
		OwnerID: userUUID,
	}

	err = h.Db.DeleteTaskByID(c.Request().Context(), params)
	if err != nil {
		return c.String(http.StatusBadRequest, "couldn't find the task")
	}

	return c.String(http.StatusNoContent, "task deleted")
}
