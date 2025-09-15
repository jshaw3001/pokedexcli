package pokeapi

import (
	"time"
	"net/http"

	"github.com/jshaw3001/pokedexcli/internal/pokecache"
)

// Client is a client for the PokeAPI.
type Client struct {
	cache	  pokecache.Cache
	httpClient *http.Client
}

// NewClient creates a new PokeAPI client with the given timeout.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}