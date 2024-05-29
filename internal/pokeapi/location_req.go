package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationResponse, error) {
	var locationResult LocationResponse
	url := baseURL + "/location"
	if pageURL != nil {
		url = *pageURL
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

	if err := json.Unmarshal(body, &locationResult); err != nil {		
		return locationResult, fmt.Errorf("could not unmarshal json %w", err)
	}
	return locationResult, nil
}