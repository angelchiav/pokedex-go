package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/angelchiav/pokedex-go/internal/pokecache"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string)
}

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = pokecache.NewCache(5 * time.Second)

func buildRegistry(command string) cliCommand {
	switch command {
	case "exit":
		return cliCommand{
			Name:        "exit",
			Description: "Exit",
			Callback:    commandExit,
		}

	case "help":
		return cliCommand{
			Name:        "help",
			Description: "Help",
			Callback:    commandHelp,
		}

	case "map":
		return cliCommand{
			Name:        "map",
			Description: "Map",
			Callback:    commandMap,
		}

	case "mapb":
		return cliCommand{
			Name:        "mapb",
			Description: "MapB",
			Callback:    commandMapb,
		}

	case "explore":
		return cliCommand{
			Name:        "explore",
			Description: "Explore",
			Callback:    commandExplore,
		}
	default:
		return cliCommand{}
	}
}

func commandExit(_ *Config, _ ...string) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp(_ *Config, _ ...string) {
	fmt.Print(`
Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
map:  Show 20 locations
mapb: Show the previous 20 locations
explore <area>: Show pokemons for a location area

`)
}

func fetchLocationAreaPage(url string) (Config, error) {

	if b, ok := cache.Get(url); ok {
		var page Config
		if err := json.Unmarshal(b, &page); err == nil {
			return page, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return Config{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Config{}, err
	}

	if res.StatusCode > 299 {
		return Config{}, fmt.Errorf("bad status %d\n%s", res.StatusCode, body)
	}

	var page Config

	if err := json.Unmarshal(body, &page); err != nil {
		return Config{}, err
	}

	cache.Add(url, body)

	return page, nil
}

type namedResource struct {
	Name string `json:"name"`
}

type pokemonEncounter struct {
	Pokemon namedResource `json:"pokemon"`
}

type locationAreaDetail struct {
	Name              string             `json:"name"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

func fetchLocationAreaDetail(url string) (locationAreaDetail, error) {
	if b, ok := cache.Get(url); ok {
		var d locationAreaDetail
		if err := json.Unmarshal(b, &d); err == nil {
			return d, nil
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreaDetail{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaDetail{}, err
	}
	if res.StatusCode > 299 {
		return locationAreaDetail{}, fmt.Errorf("bad status %d\n%s", res.StatusCode, body)
	}
	var d locationAreaDetail
	if err := json.Unmarshal(body, &d); err != nil {
		return locationAreaDetail{}, err
	}
	cache.Add(url, body)

	return d, nil
}

func commandMap(cfg *Config, _ ...string) {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	page, err := fetchLocationAreaPage(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = page.Next
	cfg.Previous = page.Previous
}

func commandMapb(cfg *Config, _ ...string) {
	if cfg.Previous == "" {
		fmt.Println("No previous page")
		return
	}

	page, err := fetchLocationAreaPage(cfg.Previous)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range page.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = page.Next
	cfg.Previous = page.Previous
}

func commandExplore(cfg *Config, args ...string) {
	if len(args) < 1 {
		fmt.Println("usage: explore <area_name>")
		return
	}

	area := strings.ToLower(args[0])
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)
	d, err := fetchLocationAreaDetail(url)
	if err != nil {
		log.Fatal(err)
	}

	if len(d.PokemonEncounters) == 0 {
		fmt.Printf("No Pokémon found in %s\n", d.Name)
		return
	}

	names := make([]string, 0, len(d.PokemonEncounters))
	for _, pe := range d.PokemonEncounters {
		names = append(names, pe.Pokemon.Name)
	}
	sort.Strings(names)
	fmt.Printf("Pokémon in %s (%d):\n", d.Name, len(names))
	for _, n := range names {
		fmt.Println(" -", n)
	}
}
