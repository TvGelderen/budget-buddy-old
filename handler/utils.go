package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/TvGelderen/budget-buddy/database"
	"github.com/TvGelderen/budget-buddy/model"
	"github.com/TvGelderen/budget-buddy/utils"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type ApiConfig struct {
    DB *database.Queries
}

func mapDbUserToUser(dbUser database.User) model.User {
    return model.User{
        Id: dbUser.ID,
        Username: dbUser.Username,
        Email: dbUser.Email,
    }
}

func mapDbTransactionToTransaction(dbTransaction database.Transaction, date time.Time) model.Transaction {
    return model.Transaction{
        Id: dbTransaction.ID,
        Amount: dbTransaction.Amount,
        Incoming: dbTransaction.Incoming,
        Description: dbTransaction.Description,
        Recurring: dbTransaction.Recurring,
        StartDate: dbTransaction.StartDate.Time,
        EndDate: dbTransaction.EndDate.Time,
        Date: date,
    }
}

func (apiCfg *ApiConfig) GetUser(r *http.Request) (model.User, error) {
    token, err := utils.GetToken(r)
    if err != nil {
        return model.User{}, err
    }

    id, err := utils.GetIdFromJWT(token)
    if err != nil {
        return model.User{}, err
    }

    user, err := apiCfg.DB.GetUserById(r.Context(), id)
    if err != nil {
        return model.User{}, err
    }

    return mapDbUserToUser(user), nil
}

func render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response())
}

func errorHTML(text string) string {
    return fmt.Sprintf("<p class='mt-6 text-error'>%s</p>", text)
}

func successHTML(text string) string {
    return fmt.Sprintf("<p class='mt-6 text-success'>%s</p>", text)
}
