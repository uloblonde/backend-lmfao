package routes

import (
	"backendtask/handlers"
	"backendtask/pkg/middleware"
	"backendtask/pkg/mysql"
	"backendtask/repositories"

	"github.com/labstack/echo/v4"
)

func AuthRout(e *echo.Group) {
	authRepo := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepo)

	e.POST("/register", h.Register)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
	e.POST("/login", h.Login)
}
