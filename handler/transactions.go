package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/TvGelderen/budget-buddy/view/transaction"
	"github.com/labstack/echo/v4"
)

func (apiCfg *ApiConfig) HandleCreateTransaction(c echo.Context) error {
	type parameters struct {
		Amount      string `json:"amount"`
		Incoming    string `json:"incoming"`
		Description string `json:"description"`
		Recurring   string `json:"recurring"`
		StartDate   string `json:"startdate"`
		EndDate     string `json:"enddate"`
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

	var endDate time.Time
	var endDateErr error

	if params.EndDate == "" {
		endDate = startDate
		endDateErr = startDateErr
	} else {
		endDate, endDateErr = time.Parse(timeFormat, params.EndDate)
	}

	_, err = apiCfg.DB.CreateTransaction(c.Request().Context(), database.CreateTransactionParams{
		UserID:      user.Id,
		Amount:      amount,
		Incoming:    incoming,
		Description: params.Description,
		Recurring:   params.Recurring,
		StartDate: sql.NullTime{
			Time:  startDate,
			Valid: startDateErr == nil,
		},
		EndDate: sql.NullTime{
			Time:  endDate,
			Valid: endDateErr == nil,
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
	}

	return c.HTML(http.StatusOK, successHTML("Transaction added successfully."))
}

func (apiCfg *ApiConfig) HandleUpdateTransactions(c echo.Context) error {
	type parameters struct {
		Id          string `json:"id"`
		Amount      string `json:"amount"`
		Incoming    string `json:"incoming"`
		Description string `json:"description"`
		Recurring   string `json:"recurring"`
		StartDate   string `json:"startdate"`
		EndDate     string `json:"enddate"`
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

	var endDate time.Time
	var endDateErr error

	if params.EndDate == "" {
		endDate = startDate
		endDateErr = startDateErr
	} else {
		endDate, endDateErr = time.Parse(timeFormat, params.EndDate)
	}

	id, err := strconv.ParseInt(params.Id, 10, 32)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.HTML(http.StatusBadRequest, errorHTML("Invalid value for id"))
	}

	err = apiCfg.DB.UpdateTransaction(c.Request().Context(), database.UpdateTransactionParams{
		ID:          int32(id),
		UserID:      user.Id,
		Amount:      amount,
		Incoming:    incoming,
		Description: params.Description,
		Recurring:   params.Recurring,
		StartDate: sql.NullTime{
			Time:  startDate,
			Valid: startDateErr == nil,
		},
		EndDate: sql.NullTime{
			Time:  endDate,
			Valid: endDateErr == nil,
		},
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
	}

	return c.HTML(http.StatusOK, successHTML("Transaction updated successfully."))
}

func (apiCfg *ApiConfig) HandleGetTransactions(c echo.Context) error {
	user := apiCfg.GetUser(c.Request())
	if user.Username == "" {
		return c.HTML(http.StatusBadRequest, errorHTML("You are not logged in."))
	}

	month := c.QueryParam("month")
	date, err := time.Parse("2006-01-02", month)

	dbTransactions, err := apiCfg.DB.GetUserTransactionsByMonth(c.Request().Context(), database.GetUserTransactionsByMonthParams{
		UserID: user.Id,
		StartDate: sql.NullTime{
			Time:  date.AddDate(0, 1, 0),
			Valid: err == nil,
		},
		EndDate: sql.NullTime{
			Time:  date,
			Valid: err == nil,
		},
	})
	if err != nil {
		return c.HTML(http.StatusBadRequest, errorHTML("Something went wrong."))
	}

	var transactions []model.Transaction

	for i := 0; i < len(dbTransactions); i++ {
		var transactionDate = dbTransactions[i].StartDate

		if dbTransactions[i].Recurring == "monthly" {
			for transactionDate.Time.Before(date) {
				transactionDate.Time = transactionDate.Time.AddDate(0, 1, 0)
			}
		}
		if dbTransactions[i].Recurring == "weekly" {
			for transactionDate.Time.Before(date) {
				transactionDate.Time = transactionDate.Time.AddDate(0, 0, 7)
			}
			for transactionDate.Time.Before(date.AddDate(0, 1, 0)) &&
				transactionDate.Time.Before(dbTransactions[i].EndDate.Time) {
				transactions = append(transactions, mapDbTransactionToTransaction(dbTransactions[i], transactionDate.Time))
				transactionDate.Time = transactionDate.Time.AddDate(0, 0, 7)
			}
			continue
		}

		transactions = append(transactions, mapDbTransactionToTransaction(dbTransactions[i], transactionDate.Time))
	}

	var income float64
	var expense float64

	for i := 0; i < len(transactions); i++ {
		if transactions[i].Incoming {
			income += transactions[i].Amount
		} else {
			expense += transactions[i].Amount
		}
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date.Before(transactions[j].Date)
	})

	return render(c, transaction.Table(transactions, income, expense))
}

func (apiCfg *ApiConfig) HandleGetTransaction(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return err
	}

	transaction, err := apiCfg.DB.GetTransaction(c.Request().Context(), int32(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transaction)
}
