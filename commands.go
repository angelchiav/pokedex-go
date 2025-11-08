package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config)
}

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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
	default:
		return cliCommand{}
	}
}

func commandExit(_ *Config) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp(_ *Config) {
	fmt.Print(`
Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
map:  Show 20 locations
mapb: Show the previous 20 locations

`)
}

func fetchLocationAreaPage(url string) (Config, error) {

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
	return page, nil
}

func commandMap(cfg *Config) {
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

func commandMapb(cfg *Config) {
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
