package main

import (
	"fmt"
	"os"
)

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

	default:
		return cliCommand{}
	}
}

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}

func commandHelp() {
	fmt.Println(`
Welcome to the Pokedex!
Usage:
	
help: Displays a help message
exit: Exit the Pokedex
`)
}
