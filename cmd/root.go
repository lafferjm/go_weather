package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_weather",
	Short: "go_weather is a quick way to check weather",
	Long: `A Fast and easy way to check the weather using
        open-meto at https://open-meteo.com and created by lafferjm`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
