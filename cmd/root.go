package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Url    string
	Method string
	Count  int64
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "ptest",
	Short: "a lightweight Http benchmarking tool for testing performance",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&Url, "url", "u", "https://google", "Website url",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Method, "method", "m", "GET", "Http request method",
	)

	rootCmd.PersistentFlags().Int64VarP(
		&Count, "count", "c", 1, "Count iterations",
	)

	rootCmd.AddCommand(runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
