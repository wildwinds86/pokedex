package main

import (
	"strings"
)

func main() {

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
