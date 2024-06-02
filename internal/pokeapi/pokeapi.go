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
	bag map[string]PokemonResponse
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

type PokemonResponse struct {
	Name          string `json:"name"`
	BaseExperience int `json:"base_experience"`	
	Height    int `json:"height"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`		
	} `json:"moves"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}



	
	