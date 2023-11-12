package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/lafferjm/go_weather/weather"
	"github.com/spf13/cobra"
)

var City string

func init() {
	forecastCmd.Flags().StringVarP(&City, "city", "c", "", "City to search weather for")
	forecastCmd.MarkFlagRequired("city")

	rootCmd.AddCommand(forecastCmd)
}

func getChoices(locations []weather.Location) []string {
	var choices []string
	for _, location := range locations {
		locationString := fmt.Sprintf("%s, %s", location.City, location.State)
		choices = append(choices, locationString)
	}

	return choices
}

func displayForecast(forecast weather.Forecast) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "Day\tMin Temp \tMax Temp")
	for i := 0; i < len(forecast.Daily.Time); i++ {
		time := forecast.Daily.Time[i]
		minTemp := forecast.Daily.MinTemperature[i]
		maxTemp := forecast.Daily.MaxTemperature[i]
		fmt.Fprintf(w, "%s\t%0.1f\t%0.1f\n", time, minTemp, maxTemp)
	}
	w.Flush()
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Checks the forecast for [city]",
	Run: func(cmd *cobra.Command, args []string) {
		cities, err := weather.SearchLocation(City)
		if err != nil {
			fmt.Printf("Error looking up forecast for: %s\n", City)
			os.Exit(1)
		}

		if len(cities.Results) == 0 {
			fmt.Println("Location not found")
			os.Exit(1)
		}

		choiceIndex := 0
		if len(cities.Results) > 1 {
			choices := getChoices(cities.Results)
			prompt := &survey.Select{
				Message: "Please choose a location:",
				Options: choices,
			}
			survey.AskOne(prompt, &choiceIndex)
		}

		forecast, err := weather.GetForecast(cities.Results[choiceIndex])
		if err != nil {
			fmt.Printf("Error looking up forecast for :%s\n", City)
			os.Exit(1)
		}

		displayForecast(*forecast)
	},
}
