package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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
