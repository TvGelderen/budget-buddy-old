package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/utils"
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
        return c.HTML(http.StatusBadRequest, "<strong class='mt-4 text-red'>Something went wrong.</strong>")
    }

    passwordHash, err := utils.HashPassword(params.Password)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, "<strong class='mt-4 text-red'>Something went wrong.</strong>")
    }

    _, err = apiCfg.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
        Name: params.Name,
        Email: params.Email,
        PasswordHash: passwordHash,
    })
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusInternalServerError, "<strong class='mt-4 text-red'>Something went wrong.</strong>")
    }

    return c.HTML(http.StatusOK, "<strong class='mt-4'>Account successfully created.</strong>")
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
        return c.HTML(http.StatusBadRequest, "<strong class='mt-4 text-red'>Something went wrong.</strong>")
    }

    user, err := apiCfg.DB.GetUserByEmail(c.Request().Context(), params.Email)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusUnauthorized, "<strong class='mt-4 text-red'>Wrong email or password.</strong>")
    }

    validPassword := utils.CheckPasswordWithHash(params.Password, user.PasswordHash)
    if !validPassword {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusUnauthorized, "<strong class='mt-4 text-red'>Wrong email or password.</strong>")
    }

    token, err := utils.CreateNewJWT(user.ID, user.Name)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusInternalServerError, "<strong class='mt-4 text-red'>Something went wrong.</strong>")
    }

    utils.SetToken(c.Response().Writer, token)

    return c.HTML(http.StatusOK, "<strong class='mt-4'>Loggin in...</strong>")
}
