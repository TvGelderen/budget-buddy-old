package handler

import (
	"fmt"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ApiConfig struct {
    DB *database.Queries
}

func GetUser() model.User {
    userDto := model.User{
        Name: "Tester01",
        Email: "test@test.com",
    }

    return userDto
}

func render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response())
}

func errorHTML(text string) string {
    return fmt.Sprintf("<p class='mt-4 text-error'>%s</p>", text)
}

func successHTML(text string) string {
    return fmt.Sprintf("<p class='mt-4 text-success'>%s</p>", text)
}
