package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// List locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		// cache hit
		//fmt.Println("Cache hit!!")
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(dat, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return RespShallowLocations{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	//fmt.Println("....Adding to cache")
	c.cache.Add(url, dat)

	return locationsResp, nil

}

// List Pokemon in specified area
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName

	// check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		// cache hit
		//fmt.Println("Cache hit!!")
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(dat, &locationAreaResp)
		if err != nil {
			return LocationArea{}, err
		}

		return locationAreaResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationAreaResp := LocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}
	//fmt.Println("....Adding to cache")
	c.cache.Add(url, dat)

	return locationAreaResp, nil

}
