package main

import (
	"fmt"
	"net/http"
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

    fs := http.FileServer(http.Dir("assets"))
    app.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    app.GET("/", userHandler.HandleHomePage)
    app.GET("/dashboard", userHandler.HandleDashboardPage)
    app.GET("/user", userHandler.HandleUserShow)

    app.GET("/login", userHandler.HandleLoginPage)
    app.GET("/register", userHandler.HandleLoginPage)
     
    app.POST("/api/login", userHandler.HandleLogin)

    app.Start(":" + port)
}
