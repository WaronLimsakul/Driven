package handlers

import (
	"database/sql"
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePutTaskKeys(c echo.Context) error {
	userIDStr := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.Logger().Printf("at HandlePutTaskKeys: couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	taskID := c.Param("id")
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		c.Logger().Printf("at HandlePutTaskKeys: couldn't parse task id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	inputKeys := c.FormValue("task-keys")
	if len(inputKeys) == 0 {
		return c.String(http.StatusBadRequest, "please provide some text")
	}

	updateTaskKeysParams := database.UpdateTaskKeysParams{
		Keys:    sql.NullString{String: inputKeys, Valid: true},
		OwnerID: userUUID,
		ID:      taskUUID,
	}

	_, err = h.Db.UpdateTaskKeys(c.Request().Context(), updateTaskKeysParams)
	if err != nil {
		return c.String(http.StatusUnauthorized, "user not an owner of the task")
	}

	return c.String(http.StatusCreated, "saved!")
}
