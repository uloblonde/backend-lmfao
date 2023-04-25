package routes

import (
	"backendtask/handlers"
	"backendtask/pkg/middleware"
	"backendtask/pkg/mysql"
	"backendtask/repositories"

	"github.com/labstack/echo/v4"
)

func TransRoute(e *echo.Group) {
	transRepo := repositories.RepositoryTrans(mysql.DB)
	h := handlers.HandlersTrans(transRepo)

	e.GET("/caritrans", middleware.Auth(h.CariTrans))
	e.GET("/transaction/:id", middleware.Auth(h.DapatTrans))
	e.POST("/membuattrans", middleware.Auth(h.MembuatTrans))
	e.POST("/notification", h.Notification)
	e.DELETE("/transaction/:id", middleware.Auth(h.HapusTrans))
}
