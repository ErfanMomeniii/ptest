package cmd

import (
	"github.com/ErfanMomeniii/colorful"
	"github.com/enescakir/emoji"
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
	})
}
func runFunc(_ *cobra.Command, _ []string) {

}
