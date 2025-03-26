package handlers

import (
	"net/http"

	"github.com/WaronLimsakul/Driven/internal/templates"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h DBHandler) HandleGetHome(c echo.Context) error {
	// will this be costly? I get user everytime
	userIDString := c.Request().Header.Get("Driven-userID")
	isSignedIn := false
	userName := ""
	if userIDString != "" {
		isSignedIn = true
		userUUID, err := uuid.Parse(userIDString)
		if err != nil {
			c.Logger().Errorf("at HandleGetHome: couldn't parse user uuid: %v", err)
			return c.Redirect(http.StatusSeeOther, "/error")
		}

		user, err := h.Db.GetUserByID(c.Request().Context(), userUUID)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/signin")
		}
		userName = user.Name
	}

	page := templates.Home()
	script := templates.HomePageScript()
	layout := templates.InjectedAppLayout
	return render(http.StatusOK, c, layout(page, script, isSignedIn, userName))
}
