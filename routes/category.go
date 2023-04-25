package routes

import (
	"backendtask/handlers"
	"backendtask/pkg/middleware"
	"backendtask/pkg/mysql"
	"backendtask/repositories"

	"github.com/labstack/echo/v4"
)

func CatRoute(e *echo.Group) {
	catRepo := repositories.RepositoryCat(mysql.DB)
	h := handlers.HandlersCat(catRepo)

	e.GET("/caricat", h.CariCat)
	e.GET("/category/:id", middleware.Auth(h.DapatCat))
	e.POST("/createcat", middleware.Auth(h.MembuatCat))
	e.PATCH("/cat/:id", middleware.Auth(h.UpdateCat))
	e.DELETE("/cat/:id", middleware.Auth(h.HapusCat))
}
