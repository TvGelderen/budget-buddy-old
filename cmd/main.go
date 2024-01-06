package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TvGelderen/budget-buddy/database"
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
    
    dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
    if dbConnectionString == "" {
        log.Fatal("No database connection string found.")
    }

    connection, err := sql.Open("postgres", dbConnectionString)
    if err != nil {
        log.Fatal("Unable to establish connection with database: ", err)
    }

    apiCfg := handler.ApiConfig {
        DB: database.New(connection),
    }
    
    app := echo.New()

    fs := http.FileServer(http.Dir("assets"))
    app.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", fs)))

    app.GET("/", apiCfg.HandleHomePage)
    app.GET("/dashboard", apiCfg.HandleDashboardPage)
    app.GET("/user", apiCfg.HandleUserShow)

    app.GET("/login", apiCfg.HandleLoginPage)
    app.GET("/register", apiCfg.HandleRegisterPage)
     
    app.POST("/api/login", apiCfg.HandleLogin)
    app.POST("/api/register", apiCfg.HandleRegister)

    app.Start(":" + port)
}
