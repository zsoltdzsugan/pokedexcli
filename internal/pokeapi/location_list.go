package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedVal, exists := c.cache.Get(url)
	if exists {

		locationRes := RespShallowLocations{}
		err := json.Unmarshal(cachedVal, &locationRes)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationRes, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationRes := RespShallowLocations{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)
	return locationRes, nil
}

func (c *Client) ListFoundPokemons(pageURL *string, location string) (RespLocation, error) {
	url := baseUrl + "/location-area/" + location
	if pageURL != nil {
		parts := strings.Split(*pageURL, "?")
		if len(parts) > 1 {
			url = parts[0] + "/" + location
		}
	}

	cachedVal, exists := c.cache.Get(url)
	if exists {

		locationRes := RespLocation{}
		err := json.Unmarshal(cachedVal, &locationRes)
		if err != nil {
			return RespLocation{}, err
		}

		return locationRes, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationRes := RespLocation{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return RespLocation{}, err
	}

	return locationRes, nil
}
