package main

import (
	"fmt"
	"os"

	"github.com/TvGelderen/budget-buddy/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
    godotenv.Load(".env")

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
        fmt.Print("PORT is missing, defaulting to 3000")
    }

    app := echo.New()

    userHandler := handler.UserHandler{}

    app.GET("/", userHandler.HandleHomePageShow)
    app.GET("/dashboard", userHandler.HandleDashboardPageShow)
    app.GET("/user", userHandler.HandleUserShow)

    app.Start(":" + port)

    fmt.Print("Hello world\n")
}
