package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PokemonData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

func ValidatePokemons(pokemons []string) ([]PokemonData, error) {
	var validPokemons []PokemonData
	var invalidPokemons []string

	for _, pokemon := range pokemons {
		if pokemon == "" {
			return nil, errors.New("pokemon name cannot be empty")
		}

		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error fetching data from pokeapi.co for %s: %w", pokemon, err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			invalidPokemons = append(invalidPokemons, pokemon)
			continue
		}

		var pokemonData PokemonData
		if err := json.NewDecoder(res.Body).Decode(&pokemonData); err != nil {
			return nil, fmt.Errorf("error decoding data from pokeapi.co for %s: %w", pokemon, err)
		}

		validPokemons = append(validPokemons, pokemonData)
	}

	if len(invalidPokemons) > 0 {
		var errorMsg string
		if len(invalidPokemons) > 1 {
			errorMsg = "Invalid pokemons: "
		} else {
			errorMsg = "Invalid pokemon: "
		}
		for i := range invalidPokemons {
			if i > 0 {
				errorMsg += ", "
			}
			errorMsg += invalidPokemons[i]
		}
		return validPokemons, errors.New(errorMsg)
	}

	return validPokemons, nil
}
