package pokeapi

import (
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
