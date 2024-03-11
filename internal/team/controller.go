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
