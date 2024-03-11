package team

import (
	"encoding/json"
	"errors"

	"github.com/mikaellpc4/go-pokemon-challenge/initializers"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/pokemon"
)

type Team struct {
	ID       int       `json:"id"`
	Owner    string    `json:"owner"`
	Pokemons []Pokemon `json:"pokemons"`
}

type Pokemon struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func GetTeamsService() (map[int]Team, error) {
	db := initializers.DB()
	defer db.Close()

	rows, err := db.Query("SELECT id, owner, pokemons FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := make(map[int]Team)
	id := 1
	for rows.Next() {
		team := Team{}
		var pokemonsJSON string
		err = rows.Scan(&team.ID, &team.Owner, &pokemonsJSON)
		if err != nil {
			return nil, err
		}

		var pokemons []Pokemon
		if err := json.Unmarshal([]byte(pokemonsJSON), &pokemons); err != nil {
			return nil, err
		}

		team.Pokemons = pokemons
		teams[id] = team
		id++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teams, nil
}

func GetTeamByOwnerService(owner string) (*Team, error) {
	db := initializers.DB()
	defer db.Close()

	var team Team

	var pokemonsJSON string
	err := db.QueryRow("SELECT * FROM teams WHERE owner = $1", owner).Scan(&team.ID, &team.Owner, &pokemonsJSON)
	if err != nil {
		return nil, err
	}

	var pokemons []Pokemon
	err = json.Unmarshal([]byte(pokemonsJSON), &pokemons)
	if err != nil {
		return nil, err
	}

	team.Pokemons = pokemons

	return &team, nil
}

func CreateTeamService(team CreateTeamRequest) error {
	db := initializers.DB()
	defer db.Close()

	var ownerExists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM teams WHERE owner = $1)", team.User).Scan(&ownerExists)
	if err != nil {
		return err
	}
	if ownerExists {
		return errors.New("user already in use")
	}

	pokemonsData, err := pokemon.ValidatePokemons(team.Team)
	if err != nil {
		return err
	}

	pokemonsJSON, err := json.Marshal(pokemonsData)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO teams (owner, pokemons) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(team.User, string(pokemonsJSON))
	if err != nil {
		return err
	}

	return nil
}
