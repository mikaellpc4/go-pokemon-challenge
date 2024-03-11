package team

import (
	"errors"

	"github.com/lib/pq"
	"github.com/mikaellpc4/go-pokemon-challenge/initializers"
	"github.com/mikaellpc4/go-pokemon-challenge/internal/pokemon"
)

func CreateTeamService(name string, pokemons []string) error {
	db := initializers.DB()
	defer db.Close()

	var userExists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM teams WHERE name = $1)", name).Scan(&userExists)
	if err != nil {
		return err
	}
	if userExists {
		return errors.New("user already in use")
	}

	err = pokemon.ValidatePokemons(pokemons)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO teams (name, pokemons) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	pokemonsArray := pq.Array(pokemons)

	_, err = stmt.Exec(name, pokemonsArray)
	if err != nil {
		return err
	}

	return nil
}
