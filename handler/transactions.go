package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/TvGelderen/budget-buddy/view/components"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleCreateTransaction(c echo.Context) error {
    type parameters struct {
        Amount string `json:"amount"`
        Incoming string `json:"incoming"`
        Description string `json:"description"`
        Recurring string `json:"recurring"`
        StartDate string `json:"startdate"`
        EndDate string `json:"enddate"`
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

    timeFormat := "2006-01-02"

    incoming := params.Incoming != "0"
    startDate, startDateErr := time.Parse(timeFormat, params.StartDate)
    endDate, endDateErr := time.Parse(timeFormat, params.EndDate)

    _, err = apiCfg.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
        UserID: user.Id,
        Amount: amount,
        Incoming: incoming,
        Description: params.Description,
        Recurring: params.Recurring,
        StartDate: sql.NullTime{
            Time: startDate,
            Valid: startDateErr == nil,
        },
        EndDate: sql.NullTime{
            Time: endDate,
            Valid: endDateErr == nil,
        },
    })
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    return c.HTML(http.StatusOK, successHTML("Transaction added successfully."))
}

func (apiCfg *ApiConfig) HandleGetTransactions(c echo.Context) error {
    user := apiCfg.GetUser(c.Request())
    if user.Username == "" {
        return c.HTML(http.StatusBadRequest, errorHTML("You are not logged in."))
    }

    month := c.QueryParam("month")
    date, err := time.Parse("2006-01-02", month)

    fmt.Printf("%v\n", date)

    dbTransactions, err := apiCfg.DB.GetUserTransactionsByMonth(c.Request().Context(), database.GetUserTransactionsByMonthParams{
        UserID: user.Id,
        StartDate: sql.NullTime{
            Time: date,
            Valid: err == nil,
        },
    })
    if err != nil {
        return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
    }

    var transactions []model.Transaction

    for i := 0; i < len(dbTransactions); i++ {
        transactions = append(transactions, mapDbTransactionToTransaction(dbTransactions[i]))
    }
    
    return render(c, components.TransactionsTable(transactions))
}
