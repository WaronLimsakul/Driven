package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostTask(c echo.Context) error {
	userID := c.Request().Header.Get("Driven-userID")
	if userID == "" {
		c.Logger().Errorf("couldn't find user id even after auth middleware")
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Errorf("couldn't parse user id: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	taskName := c.FormValue("task-name")

	priority := c.FormValue("task-priority")
	taskPriority, err := strconv.Atoi(priority)
	if err != nil {
		c.Logger().Errorf("couldn't parse task priority: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	if taskPriority < 0 || taskPriority > 3 {
		return c.String(http.StatusForbidden, "invalid priority value")
	}

	date := c.FormValue("task-date")
	// We actually expect the client to send UTC time
	taskDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		c.Logger().Errorf("couldn't parse task date: %v", err)
		return c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	now := time.Now().UTC()
	// today time at midnight
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	moreThanYear := taskDate.After(today.AddDate(1, 0, 0))
	inThePast := taskDate.Before(today)
	// don't allow any task later than a year
	if moreThanYear || inThePast {
		// fmt.Printf("than year: %v | past: %v\n", moreThanYear, inThePast)
		// fmt.Printf("today: %v\n", today)
		// fmt.Printf("task date: %v\n", taskDate)
		return c.String(http.StatusForbidden, "invalid task date")
	}

	createTaskParams := database.CreateTaskParams{
		ID:       uuid.New(),
		OwnerID:  userUUID,
		Name:     taskName,
		Priority: int32(taskPriority),
		Date:     taskDate,
	}

	_, err = h.Db.CreateTask(c.Request().Context(), createTaskParams)
	if err != nil {
		c.Logger().Errorf("couldn't create task: %v", err)
		msg := fmt.Sprintf("error: %s", err)
		return c.String(http.StatusInternalServerError, msg)
	}

	return c.String(http.StatusCreated, "create task success")
}
