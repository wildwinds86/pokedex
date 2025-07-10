package main

const (
	promptText = "Pokedex>"
	apiBaseURL = "https://pokeapi.co/api/v2"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type RespLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type config struct {
	pokeapiClient    Client
	nextLocationsURL *string
	prevLocationsURL *string
}
