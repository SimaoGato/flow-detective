package main

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flow",
	Short: "Flow Detective: Track your dev time without losing flow.",
	Long: `Flow Detective is a CLI sidecar that listens to your git commits 
and helps you track time against your sprint stories automatically.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
