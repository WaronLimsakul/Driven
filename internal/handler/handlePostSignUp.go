package handlers

import (
	"net/http"
	"time"

	"github.com/WaronLimsakul/Driven/internal/auth"
	"github.com/WaronLimsakul/Driven/internal/database"
	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandlePostSignUp(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := h.Db.GetUserByEmail(c.Request().Context(), email)

	// there is existed email
	if err == nil {
		return c.String(http.StatusConflict, "*Email already existed")
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot hash password")
	}

	createUserParams := database.CreateUserParams{
		ID:             uuid.New(),
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
	}

	user, err := h.Db.CreateUser(c.Request().Context(), createUserParams)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot create new user")
	}

	accessToken, refreshToken, err := createDoubleTokens(c, user.ID, h.JWTSecret)
	if err != nil {
		return err
	}

	AssignAuthCookies(c, h.Env == "production", accessToken, refreshToken)
	refreshTokenParams := database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiredAt: time.Now().Add(refreshExpireTime),
	}

	_, err = h.Db.CreateRefreshToken(c.Request().Context(), refreshTokenParams)
	if err != nil {
		c.Logger().Errorf("At handlePostSignUp, cannot add refresh token to db: %v", err)
		return c.String(500, "Something wen wrong")
	}

	return render(http.StatusCreated, c, templates.SignUpSuccessMessage())
}
