package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "dota2-counterpick-cli-tool",
	Long: `Dota 2 Counter Pick Is A Simple CLI Tool That Will Help You Counter Pick A Hero Easily.
-------------------------------------------------------
Single Hero Command Form:
- d2counterpick counter 'hero name'
Multiple Heroes Command Form:
- d2counterpick counter [-m | --multiple ] 'hero name 1' 'hero name 2'
Note:
Use Dashes Instead Of White-Space;
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
