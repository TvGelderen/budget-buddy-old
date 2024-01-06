package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleLogin(c echo.Context) error {
    return c.HTML(http.StatusOK, "<strong>Success!</strong>")
}
