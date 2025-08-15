package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if data, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(data, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
