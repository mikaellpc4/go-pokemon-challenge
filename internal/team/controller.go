package team

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CreateTeamRequest struct {
	User string   `json:"user" validate:"required"`
	Team []string `json:"team" validate:"required,min=1,dive,required"`
}

type ValidationError struct {
	Key   string `json:"Key"`
	Error string `json:"Error"`
}

func CreateTeam(c echo.Context) error {
	body := new(CreateTeamRequest)

	_ = c.Bind(body)

	err := c.Validate(body)
	if err != nil {
		var errors []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := err.ActualTag()
			keyType := err.Type().Name()
			keyParam := err.Param()
			if keyType != "" {
				errorMessage = errorMessage + " " + keyType
			}
			if keyParam != "" {
				errorMessage = errorMessage + " " + keyParam
			}
			validationError := ValidationError{
				Key:   err.Namespace(),
				Error: errorMessage,
			}

			errors = append(errors, validationError)
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errors})
	}

	team := CreateTeamRequest{
		User: body.User,
		Team: body.Team,
	}

	err = CreateTeamService(team)
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
