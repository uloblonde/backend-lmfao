package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	IniRoutes(e)
	AuthRout(e)
	FilmRoute(e)
	CatRoute(e)
	TransRoute(e)
	EpisodeRoute(e)
}
