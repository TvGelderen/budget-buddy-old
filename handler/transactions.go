package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleCreateTransaction(c echo.Context) error {
    type parameters struct {
        Amount string `json:"amount"`;
        Incoming string `json:"incoming"`;
        Recurring string `json:"recurring"`;
    }
    
    user := apiCfg.GetUser(c.Request())
    if user.Username == "" {
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }
    
    params := parameters{}
    err := json.NewDecoder(c.Request().Body).Decode(&params)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    amount, err := strconv.ParseFloat(params.Amount, 64)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Invalid value for amount"))
    }

    incoming := params.Recurring == "0"

    _, err = apiCfg.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
        UserID: user.Id,
        Amount: amount,
        Incoming: incoming,
        Recurring: params.Recurring,
    })
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    return c.HTML(http.StatusOK, successHTML("Transaction added successfully."))
}
