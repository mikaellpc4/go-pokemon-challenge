package team

import (
	"errors"

	"github.com/lib/pq"
	"github.com/mikaellpc4/go-pokemon-challenge/initializers"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/pokemon"
)

type Team struct {
	ID       int      `json:"id"`
	Owner    string   `json:"owner"`
	Pokemons []string `json:"pokemons"`
}

func GetTeamsService() ([]Team, error) {
	var teams []Team

	db := initializers.DB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var team Team

		pokemonsArray := pq.Array(&team.Pokemons)
		if err := rows.Scan(&team.ID, &team.Owner, pokemonsArray); err != nil {
			return nil, err
		}

		teams = append(teams, team)
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

	pokemonsArray := pq.Array(&team.Pokemons)
	err := db.QueryRow("SELECT * FROM teams WHERE owner = $1", owner).Scan(&team.ID, &team.Owner, pokemonsArray)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func CreateTeamService(owner string, pokemons []string) error {
	db := initializers.DB()
	defer db.Close()

	var ownerExists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM teams WHERE owner = $1)", owner).Scan(&ownerExists)
	if err != nil {
		return err
	}
	if ownerExists {
		return errors.New("user already in use")
	}

	err = pokemon.ValidatePokemons(pokemons)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO teams (owner, pokemons) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	pokemonsArray := pq.Array(pokemons)

	_, err = stmt.Exec(owner, pokemonsArray)
	if err != nil {
		return err
	}

	return nil
}
