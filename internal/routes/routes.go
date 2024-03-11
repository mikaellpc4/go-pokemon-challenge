package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/team"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("api/teams", team.GetTeams)
	e.GET("api/teams/:name", team.GetTeamByUser)
	e.POST("/api/teams", team.CreateTeam)

	return e
}
