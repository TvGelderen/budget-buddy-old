package middleware

import (
	"net/http"

	"github.com/TvGelderen/budget-buddy/utils"
	"github.com/labstack/echo/v4"
)

func validAuthToken(r *http.Request) error {
    token, err := utils.GetToken(r)
    if err != nil {
        return err
    }

    _, err = utils.GetIdFromJWT(token)
    if err != nil {
        return err
    }

    return nil
}

func AuthorizePage(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        if err := validAuthToken(c.Request()); err != nil {
            return c.Redirect(302, "/login")
        }

        return next(c)
    }
}
