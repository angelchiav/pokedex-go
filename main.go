package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

commands := map[string]cliCommand{
			"exit": {
				name:		 "exit",
				description: "Exit the Pokedex",
				callback:	 commandExit,	
			},

func cleanInput(text string) []string {
	re := regexp.MustCompile(`[A-Za-zÀ-ÿ]+`)
	return re.FindAllString(text, -1)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	reader := os.Stdin
	scanner := bufio.NewScanner(reader)
	for {

		fmt.Printf("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		text := cleanInput(scanner.Text())

		if len(text) == 0 {
			continue
		}
		}
	}
}
