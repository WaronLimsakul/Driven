package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(status int, c echo.Context, t templ.Component) error {
	c.Response().Status = status
	return t.Render(c.Request().Context(), c.Response())
}
