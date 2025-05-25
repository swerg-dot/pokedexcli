package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint
		
	// Check the cache first
	if dat, ok := c.cache.Get(fullURL); ok {
		// Cache hit
		//fmt.Println("cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, fmt.Errorf("cache data decode error: %v", err)
		}
		return pokemon, nil
	}

	// Cache miss, make the HTTP request
	//fmt.Println("cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Read the response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to read response body: %v", err)
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	c.cache.Add(fullURL, dat)

	return pokemon, nil
}

