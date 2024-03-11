package team

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateTeamRequest struct {
	Name     string   `json:"name"`
	Pokemons []string `json:"pokemons"`
}

func CreateTeam(c echo.Context) error {
	body := new(CreateTeamRequest)

	bindErr := c.Bind(body)
	if bindErr != nil {
		return bindErr
	}

	err := CreateTeamService(body.Name, body.Pokemons)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func GetTeams(c echo.Context) error {
	res, err := GetTeamsService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, res)
}

func GetTeamByUser(c echo.Context) error {
	name := c.Param("name")

	team, err := GetTeamByOwnerService(name)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, team)
}
