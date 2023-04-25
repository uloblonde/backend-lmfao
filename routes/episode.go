package routes

import (
	"backendtask/handlers"
	"backendtask/pkg/middleware"
	"backendtask/pkg/mysql"
	"backendtask/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoute(e *echo.Group) {
	epiRepo := repositories.RepositoryEpi(mysql.DB)
	h := handlers.HandlersEpi(epiRepo)

	e.GET("/cariepi", h.CariEpi)
	e.GET("/epi/:id", h.DapatEpi)
	e.GET("/film/:id/episode", h.CariEpiByFilm)
	e.GET("/film/:id/episode/:ide", h.DapatEpiByFilm)
	e.POST("/createepi", middleware.Auth(middleware.UploadFile(h.MembuatEpi)))
	e.PATCH("/epi/:id", middleware.Auth(middleware.UploadFile(h.UpdateEpi)))
	e.DELETE("/epi/:id", h.HapusEpi)
}
