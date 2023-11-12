package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GeocodingResponse struct {
	Results []Location `json:"results"`
}

type Location struct {
	City      string  `json:"name"`
	State     string  `json:"admin1"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func filterResults(locations []Location, city string) []Location {
	var matchedLocations []Location

	for _, location := range locations {
		if strings.HasPrefix(location.City, city) {
			matchedLocations = append(matchedLocations, location)
		}
	}

	return matchedLocations
}

func SearchLocation(city string) (*GeocodingResponse, error) {
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s", city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geocodingResponse GeocodingResponse
	err = json.NewDecoder(resp.Body).Decode(&geocodingResponse)
	if err != nil {
		return nil, err
	}

	geocodingResponse.Results = filterResults(geocodingResponse.Results, city)

	return &geocodingResponse, nil
}
