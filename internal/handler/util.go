package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(status int, context echo.Context, templComp templ.Component) error {
	context.Response().Status = status
	return templComp.Render(context.Request().Context(), context.Response())
}
