package pokeapi

import (
	"net/http"
	"time"

	"github.com/yongebai/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: time.Second * 30,
		},
		cache: pokecache.NewCache(interval),
	}
}
	
type LocationResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {	
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`		
	} `json:"pokemon_encounters"`
}