package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const promptText string = "Pokedex>"

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

//var cmds map[string]cliCommand

func startRepl() {

	cmds := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	fmt.Print(promptText)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			cmd, exists := cmds[input[0]]
			if exists {
				fmt.Print(cmd.name)
				cmd.callback()
			}
		}
		fmt.Print(promptText)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Useage:")
	return nil
}
