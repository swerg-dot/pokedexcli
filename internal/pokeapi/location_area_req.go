package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// Check the cache first
	if dat, ok := c.cache.Get(fullURL); ok {
		// Cache hit
		//fmt.Println("cache hit!")
		var locationAreasResp LocationAreasResp
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, fmt.Errorf("cache data decode error: %v", err)
		}
		return locationAreasResp, nil
	}

	// Cache miss, make the HTTP request
	//fmt.Println("cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResp{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Read the response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var locationAreasResp LocationAreasResp
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

//?

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint
		
	// Check the cache first
	if dat, ok := c.cache.Get(fullURL); ok {
		// Cache hit
		//fmt.Println("cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, fmt.Errorf("cache data decode error: %v", err)
		}
		return locationArea, nil
	}

	// Cache miss, make the HTTP request
	//fmt.Println("cache miss!")

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Read the response body
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to read response body: %v", err)
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	c.cache.Add(fullURL, dat)

	return locationArea, nil
}