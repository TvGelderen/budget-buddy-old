package handler

import (
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/TvGelderen/budget-buddy/view/user"
	"github.com/TvGelderen/budget-buddy/view/pages"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {}

func (h UserHandler) HandleUserShow(c echo.Context) error {
    userDto := model.User{
        Name: "Tester01",
        Email: "test@test.com",
    }

    return render(c, user.Show(userDto))
}

func (h UserHandler) HandleHomePage(c echo.Context) error {
    userDto := model.User{
        Name: "Tester01",
        Email: "test@test.com",
    }

    return render(c, pages.Index(userDto)); 
}

func (h UserHandler) HandleDashboardPage(c echo.Context) error {
    userDto := model.User{
        Name: "Tester01",
        Email: "test@test.com",
    }

    return render(c, pages.Dashboard(userDto)); 
}

func (h UserHandler) HandleLoginPage(c echo.Context) error {
    userDto := model.User{
        Name: "Tester01",
        Email: "test@test.com",
    }

    return render(c, pages.Login(userDto));
}
