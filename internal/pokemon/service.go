package pokemon

import (
	"errors"
	"fmt"
	"net/http"
)

func ValidatePokemons(pokemons []string) error {
	var pokemonErrors []error

	for _, pokemon := range pokemons {
		if pokemon == "" {
			return errors.New("pokemon can't have a empty string as name")
		}
		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

		res, err := http.Get(url)
		if err != nil {
			return errors.New("error when fetching data from pokeapi.co")
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			pokemonErrors = append(pokemonErrors, fmt.Errorf(pokemon))
		}
	}

	if len(pokemonErrors) > 0 {
		var errorMsg string
		if len(pokemonErrors) > 1 {
			errorMsg = "Invalid pokemons: "
		} else {
			errorMsg = "Invalid pokemon: "
		}
		for i, err := range pokemonErrors {
			if i > 0 {
				errorMsg += ", "
			}
			errorMsg += err.Error()
		}
		return errors.New(errorMsg)
	}

	return nil
}
