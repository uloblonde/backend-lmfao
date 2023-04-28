package main

import (
	"backendtask/database"
	"backendtask/pkg/mysql"
	"backendtask/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	mysql.DatabaseInit()
	database.RunMigrations()

	// e.GET("/Login", Lgn)
	// e.GET("/Passsword", Pw)

	e.Static("/uploads", "./uploads")
	routes.RouteInit(e.Group("/api/v1"))
	var PORT = os.Getenv("PORT")
	fmt.Println("server running localhost:" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT)) // delete localhost
}
