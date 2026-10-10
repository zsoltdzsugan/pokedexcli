package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseUrl + "/location-area/" + name

	if cachedVal, exists := c.cache.Get(url); exists {

		locationResp := Location{}
		err := json.Unmarshal(cachedVal, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

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
