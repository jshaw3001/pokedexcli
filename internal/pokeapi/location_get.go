package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Explore - takes a location area name as an argument and retrieves a list of Pokemon that can be encountered in that area.

func (c *Client) GetLocation(location string) (Location, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	
	return locationResp, nil
}
