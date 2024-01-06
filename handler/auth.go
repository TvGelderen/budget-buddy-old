package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleRegister(c echo.Context) error {
    type parameters struct {
        Name string `json:"name"`;
        Email string `json:"email"`;
        Password string `json:"password"`;
    }
    
    params := parameters{}
    err := json.NewDecoder(c.Request().Body).Decode(&params)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, "<strong class='text-red'>Something went wrong.</strong>")
    }

    fmt.Printf("%v\n", params)

    return c.HTML(http.StatusOK, "<strong>Success!</strong>")
}

func (apiCfg *ApiConfig) HandleLogin(c echo.Context) error {
    type parameters struct {
        Email string `json:"email"`;
        Password string `json:"password"`;
    }
    
    params := parameters{}
    err := json.NewDecoder(c.Request().Body).Decode(&params)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, "<strong class='text-red'>Something went wrong.</strong>")
    }

    fmt.Printf("%v\n", params)

    return c.HTML(http.StatusOK, "<strong>Success!</strong>")
}
