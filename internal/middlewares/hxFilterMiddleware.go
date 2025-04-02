package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// prevent user to randomly go to route that give just html component
func HXFilterMiddleWare(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		HxHeader := c.Request().Header.Get("HX-Request")
		if HxHeader != "true" {
			fmt.Printf("point 1")
			return c.Redirect(http.StatusSeeOther, "/home")
		}
		fmt.Printf("point 2")
		return handler(c)
	}
}
