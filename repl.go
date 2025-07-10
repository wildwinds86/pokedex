package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(promptText)

		scanner.Scan()
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		cmd, exists := getCommands(cfg)[input[0]]

		if exists {
			err := cmd.callback(cfg)

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command, please try again or enter 'help' to view a list of commands.")
			continue
		}
	}
}
