package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	var pokemonResult PokemonResponse
	url := baseURL+"/pokemon/"+pokemonName	

	if body, ok := c.cache.Get(url); ok {
		// fmt.Println("Cache hit")
		if err := json.Unmarshal(body, &pokemonResult); err != nil {
			return pokemonResult, fmt.Errorf("could not unmarshal json %w", err)
		}
		return pokemonResult, nil	
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {		
		return pokemonResult, fmt.Errorf("unable to get response: %w", err)
	}	
	
	resp, _ := c.httpClient.Do(req)

	if resp.StatusCode > 399 {		
		return pokemonResult, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		
		return pokemonResult, fmt.Errorf("unable to read response body: %w", err)
	}

	time.Sleep(time.Second * 2)
	// fmt.Println("Cache miss")
	c.cache.Add(url, body)

	if err := json.Unmarshal(body, &pokemonResult); err != nil {		
		return pokemonResult, fmt.Errorf("could not unmarshal json %w", err)
	}
	return pokemonResult, nil
}

