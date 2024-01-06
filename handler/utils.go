package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/TvGelderen/budget-buddy/database"
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
