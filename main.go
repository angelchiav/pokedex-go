package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func cleanInput(text string) []string {
	re := regexp.MustCompile(`[A-Za-zÀ-ÿ]+`)
	return re.FindAllString(text, -1)
}

func main() {
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)

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
			buildRegistry("exit").Callback()
			continue

		case "help":
			buildRegistry("help").Callback()
			continue

		case "map":
			buildRegistry("map").Callback()
			continue

		case "mapb":
			buildRegistry("mapb").Callback()
			continue

		default:
			fmt.Println("Unknown command")
			continue

		}
	}
}
