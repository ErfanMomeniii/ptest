package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var (
	Url     string
	Method  string
	Header  []string
	Body    string
	Timeout int64
	Count   int64
	Diagram bool
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "ptest",
	Short: "a lightweight Http benchmarking tool for testing performance",
}

func init() {
	rootCmd.PersistentFlags().BoolP(
		"help", "", false, "Help for seeing more information on other commands",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Url, "url", "u", "https://google.com", "Website url",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Method, "method", "m", "GET", "Http request method",
	)

	Header = *rootCmd.PersistentFlags().StringArrayP(
		"header", "h", []string{}, "Headers of the request",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Body, "body", "b", "", "Body for the HTTP request",
	)

	rootCmd.PersistentFlags().Int64VarP(
		&Count, "count", "c", 1, "Count iterations",
	)

	rootCmd.PersistentFlags().Int64VarP(
		&Timeout, "timeout", "t", int64(time.Second*10), "Timeout for each HTTP call",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&Diagram, "diagram", "d", false, "Should draw diagram or not",
	)

	rootCmd.Flags().SortFlags = false

	rootCmd.AddCommand(runCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
