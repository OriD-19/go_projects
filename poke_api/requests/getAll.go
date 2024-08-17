// Package requests retrieves information from the PokeAPI in JSON format
// that is going to be consumed by templates.
package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"OriD19.com/api/poke/data"
)

type result struct {
	Result []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func GetAllPaginated(page int) (*[]data.Pokemon, error) {
	const offset = 20

	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=20&offset=%d", offset*(page-1)))

	if err != nil {
		return nil, err
	}

	var ps []data.Pokemon
	var pokemonsResult result

	// Unmarshal the JSON from the main response

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &pokemonsResult)

	if err != nil {
		return nil, err
	}

	for _, c := range pokemonsResult.Result {
		n, err := ParsePokemon(c.Url)

		if err != nil {
			return nil, err
		}

		ps = append(ps, *n)
	}

	return &ps, nil
}

func ParsePokemon(pokemonUrl string) (*data.Pokemon, error) {
	res, err := http.Get(pokemonUrl)

	if err != nil {
		return nil, err
	}

	// Pokemon to return
	var p data.Pokemon

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &p)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
