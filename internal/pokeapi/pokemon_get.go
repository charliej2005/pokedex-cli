package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(url); ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(data, &pokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonRes := Pokemon{}
	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonRes, nil
}
