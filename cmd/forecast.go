package cmd

import (
	"fmt"
	"os"

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

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Checks the forecast for [city]",
	Run: func(cmd *cobra.Command, args []string) {
		forecast, err := weather.SearchLocation(City)
		if err != nil {
			fmt.Printf("Error looking up forecast for: %s\n", City)
			os.Exit(1)
		}

		if len(forecast.Results) == 0 {
			fmt.Println("Location not found")
			os.Exit(1)
		}

		choiceIndex := 0
		if len(forecast.Results) > 1 {
			choices := getChoices(forecast.Results)
			prompt := &survey.Select{
				Message: "Please choose a location:",
				Options: choices,
			}
			survey.AskOne(prompt, &choiceIndex)
		}

		fmt.Println(forecast.Results[choiceIndex])
	},
}
