package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const jwtExpireTime time.Duration = 15 * time.Minute
const refreshExpireTime time.Duration = 7 * 24 * time.Hour

func render(status int, context echo.Context, templComp templ.Component) error {
	context.Response().Status = status
	return templComp.Render(context.Request().Context(), context.Response())
}

func createDoubleTokens(c echo.Context, userID uuid.UUID, secret string) (accessToken, refreshToken string, err error) {
	accessToken, err = auth.CreateJWT(userID, jwtExpireTime, secret)
	if err != nil {
		c.Logger().Errorf("Couldn't create access token: %v", err)
		return "", "", c.String(http.StatusInternalServerError, "Couldn't create access token")
	}

	refreshToken, err = auth.CreateRefreshToken()
	if err != nil {
		c.Logger().Errorf("Couldn't create refresh token: %v", err)
		return "", "", c.String(http.StatusInternalServerError, "Couldn't create refresh token")
	}

	return accessToken, refreshToken, nil
}

func AssignAuthCookies(c echo.Context, production bool, accessToken, refreshToken string) {
	AssignAccessTokenCookie(c, accessToken, production)
	AssignRefreshTokenCookie(c, refreshToken, production)
	return
}

func AssignAccessTokenCookie(c echo.Context, accessToken string, production bool) {
	// normal browser will bind this cookie to only the server it got cookie from
	accessTokenCookie := new(http.Cookie)
	accessTokenCookie.Name = "driven-jwt"
	accessTokenCookie.Value = accessToken
	accessTokenCookie.Expires = time.Now().Add(jwtExpireTime)
	accessTokenCookie.HttpOnly = true
	if production {
		accessTokenCookie.Secure = true
	}
	accessTokenCookie.Path = "/"
	// c.SetCookie() set the response to tell client that they should include
	// the cookie next time it makes request. It's not immediate.
	c.SetCookie(accessTokenCookie)
	// So I have to set in manually only for this request
	c.Request().AddCookie(accessTokenCookie)
	return
}

func AssignRefreshTokenCookie(c echo.Context, refreshToken string, production bool) {
	refreshTokenCookie := new(http.Cookie)
	refreshTokenCookie.Name = "driven-refresh-token"
	refreshTokenCookie.Value = refreshToken
	refreshTokenCookie.Expires = time.Now().Add(refreshExpireTime)
	refreshTokenCookie.HttpOnly = true
	if production {
		refreshTokenCookie.Secure = true
	}
	c.SetCookie(refreshTokenCookie)
	return
}

func (h DBHandler) CreateTaskForUser(c echo.Context) (database.Task, error) {
	userID := c.Request().Header.Get("Driven-userID")
	if userID == "" {
		c.Logger().Errorf("couldn't find user id even after auth middleware")
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.Logger().Errorf("couldn't parse user id: %v", err)
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	taskName := c.FormValue("task-name")

	priority := c.FormValue("task-priority")
	taskPriority, err := strconv.Atoi(priority)
	if err != nil {
		c.Logger().Errorf("couldn't parse task priority: %v", err)
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong, try again later")
	}

	if taskPriority < 0 || taskPriority > 3 {
		return database.Task{}, c.String(http.StatusForbidden, "invalid priority value")
	}

	date := c.FormValue("task-date")
	// We actually expect the client to send UTC time
	taskDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		c.Logger().Errorf("couldn't parse task date: %v", err)
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong, try again later")
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
		return database.Task{}, c.String(http.StatusForbidden, "invalid task date")
	}

	createTaskParams := database.CreateTaskParams{
		ID:       uuid.New(),
		OwnerID:  userUUID,
		Name:     taskName,
		Priority: int32(taskPriority),
		Date:     taskDate,
	}

	newTask, err := h.Db.CreateTask(c.Request().Context(), createTaskParams)
	if err != nil {
		c.Logger().Errorf("couldn't create task: %v", err)
		msg := fmt.Sprintf("error: %s", err)
		return database.Task{}, c.String(http.StatusInternalServerError, msg)
	}

	return newTask, nil
}

func (h DBHandler) doneTaskForUser(c echo.Context) (database.Task, error) {
	taskIDStr := c.Param("id")
	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.Logger().Printf("couldn't parse task id: %v", err)
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong")
	}

	userIDStr := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.Logger().Printf("couldn't parse user id: %v", err)
		return database.Task{}, c.String(http.StatusInternalServerError, "something went wrong")
	}

	doneTaskArgs := database.DoneTaskByIDParams{
		ID:      taskUUID,
		OwnerID: userUUID,
	}

	// check if user is correct in one query
	updatedTask, err := h.Db.DoneTaskByID(c.Request().Context(), doneTaskArgs)
	if err != nil {
		c.Logger().Printf("couldn't update task id: %v", err)
		return database.Task{}, c.String(http.StatusUnauthorized, "something went wrong")
	}

	return updatedTask, nil
}

func (h DBHandler) undoneTaskForUser(c echo.Context) ([]database.Task, error) {
	userIDStr := c.Request().Header.Get("Driven-userID")
	userUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.Logger().Printf("at undoneTaskForUser: couldn't parse user id: %v", err)
		return []database.Task{}, c.String(http.StatusInternalServerError, "something went wrong")
	}

	taskIDStr := c.Param("id")
	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {
		c.Logger().Printf("at undoneTaskForUser: couldn't parse task id: %v", err)
		return []database.Task{}, c.String(http.StatusInternalServerError, "something went wrong")
	}

	undoneTaskParams := database.UndoneTaskByIDParams{
		ID:      taskUUID,
		OwnerID: userUUID,
	}

	updatedTask, err := h.Db.UndoneTaskByID(c.Request().Context(), undoneTaskParams)
	if err != nil {
		return []database.Task{}, c.String(http.StatusUnauthorized, "couldn't undone task: user not the owner")
	}

	// Have to render entire column again because we don't know where to put it back
	taskDate := updatedTask.Date
	getDayTasksParam := database.GetTaskByDateParams{
		OwnerID: updatedTask.OwnerID,
		Date:    taskDate,
	}
	dayTasks, err := h.Db.GetTaskByDate(c.Request().Context(), getDayTasksParam)

	if err != nil {
		c.Logger().Printf("couldn't get user tasks by date: %v", err)
		return []database.Task{}, c.String(http.StatusInternalServerError, "something went wrong")
	}

	return dayTasks, nil
}
