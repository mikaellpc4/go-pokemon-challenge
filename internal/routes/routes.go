package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/team"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.POST("/team", team.CreateTeam)

	return e
}
