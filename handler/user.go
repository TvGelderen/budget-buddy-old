package handler

import (
	"github.com/TvGelderen/budget-buddy/view/user"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleUserShow(c echo.Context) error {
    userDto := GetUser()

    return render(c, user.Show(userDto))
}
