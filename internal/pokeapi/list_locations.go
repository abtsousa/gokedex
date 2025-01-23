package pokeapi

import (
	"github.com/abtsousa/gokedex/internal/cache"
)

type LocationsList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string, cache *cache.Cache) (LocationsList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	return getData[LocationsList](c, cache, url)
}
