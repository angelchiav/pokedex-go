package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)

	cfg := &Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

	for {

		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "read error", err)
			}
			return
		}

		tokens := cleanInput(scanner.Text())
		if len(tokens) == 0 {
			continue
		}

		switch strings.ToLower(tokens[0]) {

		case "exit":
			buildRegistry("exit").Callback(cfg, tokens[1:]...)
			continue

		case "help":
			buildRegistry("help").Callback(cfg, tokens[1:]...)
			continue

		case "map":
			buildRegistry("map").Callback(cfg, tokens[1:]...)
			continue

		case "mapb":
			buildRegistry("mapb").Callback(cfg, tokens[1:]...)
			continue

		case "explore":
			buildRegistry("explore").Callback(cfg, tokens[1:]...)

		case "catch":
			buildRegistry("catch").Callback(cfg, tokens[1:]...)

		case "inspect":
			buildRegistry("inspect").Callback(cfg, tokens[1:]...)

		case "pokedex":
			buildRegistry("pokedex").Callback(cfg, tokens[1:]...)

		default:
			fmt.Println("Unknown command.  Write 'help' to learn the commands.")
			continue

		}
	}
}
