package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var City string

func init() {
	forecastCmd.Flags().StringVarP(&City, "city", "c", "", "City to search weather for")
	forecastCmd.MarkFlagRequired("city")

	rootCmd.AddCommand(forecastCmd)
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Checks the forecast for [city]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(City)
	},
}
