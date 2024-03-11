package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/team"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("api/teams", team.GetTeams)
	e.GET("api/teams/:ïd", team.CreateTeam)
	e.POST("/api/teams", team.CreateTeam)

	return e
}
