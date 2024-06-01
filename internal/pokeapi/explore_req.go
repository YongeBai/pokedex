package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ExploreLocationArea(area string) (LocationAreaResponse, error) {
	var resultLocationArea LocationAreaResponse
	if area == "" {
		return resultLocationArea, fmt.Errorf("area is nil")
	}

	url := baseURL+"/location-area/"+area
	
	if body, ok := c.cache.Get(url); ok {
		// fmt.Println("Cache hit")
		if err := json.Unmarshal(body, &resultLocationArea); err != nil {
			return resultLocationArea, fmt.Errorf("could not unmarshal json %w", err)
		}
		return resultLocationArea, nil	
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {		
		return resultLocationArea, fmt.Errorf("unable to get response: %w", err)
	}	
	
	resp, _ := c.httpClient.Do(req)

	if resp.StatusCode > 399 {		
		return resultLocationArea, fmt.Errorf("Status Code: %v, No Pokedex entry for this area", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		
		return resultLocationArea, fmt.Errorf("unable to read response body: %w", err)
	}

	time.Sleep(time.Second * 2)
	// fmt.Println("Cache miss")
	c.cache.Add(url, body)

	if err := json.Unmarshal(body, &resultLocationArea); err != nil {		
		return resultLocationArea, fmt.Errorf("could not unmarshal json %w", err)
	}
	return resultLocationArea, nil
}

