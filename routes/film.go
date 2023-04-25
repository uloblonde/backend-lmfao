package routes

import (
	"backendtask/handlers"
	"backendtask/pkg/middleware"
	"backendtask/pkg/mysql"
	"backendtask/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoute(e *echo.Group) {
	filmRepo := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlersFilm(filmRepo)

	e.GET("/CariFilm", h.CariFilm)
	e.GET("/film/:id", h.DapatFilm)
	e.POST("/buatfilm", middleware.UploadFile(h.MembuatFilm))
	e.PATCH("/film/:id", middleware.Auth(middleware.UploadFile(h.UpdateFilm)))
	e.DELETE("/film/:id", middleware.Auth(h.HapusFilm))
}
