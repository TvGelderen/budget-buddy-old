package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TvGelderen/budget-buddy/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
    godotenv.Load(".env")

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
        fmt.Print("PORT is missing, defaulting to 3000")
    }
    
    connectionString := os.Getenv("DB_CONNECTION_STRING")
    if connectionString == "" {
        log.Fatal("No database connection string found.")
    }

    app := echo.New()

    apiCfg := handler.ApiConfig{}

    fs := http.FileServer(http.Dir("assets"))
    app.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    app.GET("/", apiCfg.HandleHomePage)
    app.GET("/dashboard", apiCfg.HandleDashboardPage)
    app.GET("/user", apiCfg.HandleUserShow)

    app.GET("/login", apiCfg.HandleLoginPage)
    app.GET("/register", apiCfg.HandleLoginPage)
     
    app.POST("/api/login", apiCfg.HandleLogin)

    app.Start(":" + port)
}
