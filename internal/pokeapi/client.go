package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListLocations(pageURL *string) (LocationsList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsList{}, fmt.Errorf("couldn't make request: %v", err)
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return LocationsList{}, fmt.Errorf("couldn't get response: %v", err)
	}

	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return LocationsList{}, fmt.Errorf("couldn't parse received data: %v", err)
	}

	loc := LocationsList{}
	err = json.Unmarshal(dat, &loc)
	if err != nil {
		return LocationsList{}, fmt.Errorf("couldn't unmarshal received data: %v", err)
	}

	return loc, nil
}
