package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocations, error) {
	url := apiBaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheData, exists := pokecache.Get(url)
	if exists {
		fmt.Println("*** USING CACHED DATA ***")
		locationsResp := RespLocations{}
		err := json.Unmarshal(cacheData, &locationsResp)

		if err != nil {
			return RespLocations{}, err
		} else {
			return locationsResp, nil
		}
	}
	fmt.Println("-- No cached data found, pulling from API")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocations{}, err
	}

	locationsResp := RespLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespLocations{}, err
	}

	pokecache.Add(url, data)
	return locationsResp, nil
}

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
