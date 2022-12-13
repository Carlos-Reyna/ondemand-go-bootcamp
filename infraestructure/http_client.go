package infrastructure

import (
	"io"
	"net/http"
)

type HTTPClient interface {
	Get(url string) ([]byte, error)
}

type PokeAPIHTTPClient struct{}

func (c *PokeAPIHTTPClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
