package handler

import (
	"github.com/TvGelderen/budget-buddy/view/pages"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleHomePage(c echo.Context) error {
    userDto := GetUser()

    return render(c, pages.Index(userDto)); 
}

func (apiCfg *ApiConfig) HandleDashboardPage(c echo.Context) error {
    userDto := GetUser()

    return render(c, pages.Dashboard(userDto)); 
}

func (apiCfg *ApiConfig) HandleLoginPage(c echo.Context) error {
    userDto := GetUser()

    return render(c, pages.Login(userDto));
}
