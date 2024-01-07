package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleRegister(c echo.Context) error {
    type parameters struct {
        Username string `json:"username"`;
        Email string `json:"email"`;
        Password string `json:"password"`;
    }
    
    params := parameters{}
    err := json.NewDecoder(c.Request().Body).Decode(&params)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    passwordHash, err := utils.HashPassword(params.Password)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    _, err = apiCfg.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
        ID: uuid.New(),
        Username: params.Username,
        Email: params.Email,
        PasswordHash: passwordHash,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    })
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusInternalServerError, errorHTML("Something went wrong."))
    }

    c.Response().Writer.Header().Set("Hx-Redirect", "/")

    return nil
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
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    user, err := apiCfg.DB.GetUserByEmail(c.Request().Context(), params.Email)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusUnauthorized, errorHTML("Wrong email or password."))
    }

    validPassword := utils.CheckPasswordWithHash(params.Password, user.PasswordHash)
    if !validPassword {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusUnauthorized, errorHTML("Wrong email or password."))
    }

    token, err := utils.CreateNewJWT(user.ID, user.Username)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusInternalServerError, errorHTML("Something went wrong."))
    }

    utils.SetToken(c.Response().Writer, token)

    c.Response().Writer.Header().Set("Hx-Redirect", "/")

    return nil
}

func (apiCfg *ApiConfig) HandleLogout(c echo.Context) error {
    utils.RemoveToken(c.Response().Writer)

    return c.Redirect(302, "/")
}
