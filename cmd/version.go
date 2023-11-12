package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go_weather",
	Long:  `All software has versions.  This is go_weather's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("weather checker in golang version 0.1")
	},
}
