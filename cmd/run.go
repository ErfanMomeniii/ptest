package cmd

import (
	"github.com/ErfanMomeniii/colorful"
	"github.com/enescakir/emoji"
	"github.com/erfanmomeniii/ptest/internal/app"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tool",
	Run:   runFunc,
}

func init() {
	cobra.OnInitialize(func() {
		colorful.Printf(colorful.BlueColor, colorful.DefaultBackground, "%v  Running tool ... \n", emoji.PersonRunning)
		colorful.Printf(colorful.BlueColor, colorful.DefaultBackground, "%v  Generating result ... \n", emoji.WritingHand)
		colorful.Printf(colorful.YellowColor, colorful.DefaultBackground, "%v  CTRL+C to gracefully stop \n", emoji.Warning)
	})
}
func runFunc(_ *cobra.Command, _ []string) {
	a := app.New(Url, Method, Header, Body, Count, Timeout)

	a.Run()
}
