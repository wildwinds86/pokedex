package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const promptText string = "Pokedex>"

func main() {
	fmt.Print(promptText)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			fmt.Printf("Your command was: %s\n", input[0])
		}
		fmt.Print(promptText)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
