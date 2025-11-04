package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func()
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

func commandExit( /*config *config*/ ) {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp( /*config *config*/ ) {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
`)
}

func commandMap() {

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var maps Config

	if err := json.Unmarshal(body, &maps); err != nil {
		log.Fatal(err)
	}

	for _, loc := range maps.Results {
		fmt.Println(loc.Name)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

}

func commandMapb() {

	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	var maps Config

	if err := json.Unmarshal(body, &maps); err != nil {
		log.Fatal(err)
	}

	for _, loc := range maps.Previous.Results {
		fmt.Println(loc.Name)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

}
