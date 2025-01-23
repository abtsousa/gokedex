package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/abtsousa/gokedex/internal/cache"
	"io"
	"net/http"
)

func getData[T any](c *Client, cache *cache.Cache, url string) (T, error) {
	var t T
	if cache != nil {
		if cached, exists := cache.Get(url); exists {
			if err := json.Unmarshal(cached, &t); err == nil {
				return t, nil
			} else {
				fmt.Println("[INFO] Found invalid data in cache. Cleaning.")
				cache.Remove(url)
			}
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return t, fmt.Errorf("couldn't make request: %v", err)
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return t, fmt.Errorf("couldn't get response: %v", err)
	}

	dat, err := io.ReadAll(rsp.Body)
	if err != nil {
		return t, fmt.Errorf("couldn't parse received data: %v", err)
	}

	// Add to cache
	cache.Add(url, dat)

	err = json.Unmarshal(dat, &t)
	if err != nil {
		return t, fmt.Errorf("couldn't unmarshal received data: %v", err)
	}

	return t, nil
}
