package handler

import (
	"github.com/TvGelderen/budget-buddy/view/pages"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleHomePage(c echo.Context) error {
    user, _ := apiCfg.GetUser(c.Request())

    return render(c, pages.Index(user)); 
}

func (apiCfg *ApiConfig) HandleDashboardPage(c echo.Context) error {
    user, _ := apiCfg.GetUser(c.Request())

    return render(c, pages.Dashboard(user)); 
}

func (apiCfg *ApiConfig) HandleAnalyticsPage(c echo.Context) error {
    user, _ := apiCfg.GetUser(c.Request())

    return render(c, pages.Analytics(user)); 
}

func (apiCfg *ApiConfig) HandleRegisterPage(c echo.Context) error {
    user, err := apiCfg.GetUser(c.Request())
    if err == nil {
        return c.Redirect(302, "/")
    }

    return render(c, pages.Register(user));
}

func (apiCfg *ApiConfig) HandleLoginPage(c echo.Context) error {
    user, err := apiCfg.GetUser(c.Request())
    if err == nil {
        return c.Redirect(302, "/")
    }

    return render(c, pages.Login(user));
}
