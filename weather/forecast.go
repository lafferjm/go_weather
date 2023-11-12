package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Forecast struct {
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	Daily     DailyForecast `json:"daily"`
}

type DailyForecast struct {
	Time           []string  `json:"time"`
	MinTemperature []float64 `json:"temperature_2m_min"`
	MaxTemperature []float64 `json:"temperature_2m_max"`
}

func GetForecast(location Location) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&temperature_unit=fahrenheit&daily=temperature_2m_max,temperature_2m_min", location.Latitude, location.Longitude)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error looking up the forecast")
		os.Exit(1)
	}
	defer resp.Body.Close()

	var forecast Forecast
	err = json.NewDecoder(resp.Body).Decode(&forecast)
	if err != nil {
		fmt.Printf("Error getting forecast")
		os.Exit(1)
	}

	fmt.Println(forecast)
}
