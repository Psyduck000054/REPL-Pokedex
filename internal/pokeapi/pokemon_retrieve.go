package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PropertiesRetrieval(pokemon string) (Pokemon, error) {
	var p Pokemon

	url := baseURL + "/pokemon/" + pokemon + "/"

	// cache logic
	if val, ok := c.userCache.Get(url); ok {
		if err := json.Unmarshal(val, &p); err != nil {
			return Pokemon{}, err
		}
		return p, nil
	}

	// build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// execute it
	res, err := c.cli.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	// read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	// store in cache
	c.userCache.Add(url, body)

	// unmarshal json into go structs
	err = json.Unmarshal(body, &p)
	if err != nil {
		return Pokemon{}, err
	}

	return p, nil
}
