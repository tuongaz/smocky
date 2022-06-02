package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smocky",
	Short: "Smocky command",
	Run: func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if !showVersion {
			_ = cmd.Help()
			return
		}

		fmt.Println("Version 1.2.1")
	},
}

var version bool

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "show version")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
