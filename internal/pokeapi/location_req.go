package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetLocations(pageURL *string) (LocationResponse, error) {
	var locationResult LocationResponse
	url := baseURL+"/location?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	if body, ok := c.cache.Get(url); ok {
		fmt.Println("Cache hit")
		if err := json.Unmarshal(body, &locationResult); err != nil {
			return locationResult, fmt.Errorf("could not unmarshal json %w", err)
		}
		return locationResult, nil	
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {		
		return locationResult, fmt.Errorf("unable to get response: %w", err)
	}	
	
	resp, _ := c.httpClient.Do(req)

	if resp.StatusCode > 399 {		
		return locationResult, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		
		return locationResult, fmt.Errorf("unable to read response body: %w", err)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("Cache miss")
	c.cache.Add(url, body)

	if err := json.Unmarshal(body, &locationResult); err != nil {		
		return locationResult, fmt.Errorf("could not unmarshal json %w", err)
	}
	return locationResult, nil
}