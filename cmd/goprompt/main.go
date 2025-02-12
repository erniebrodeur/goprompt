package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is our top-level Cobra command.
var rootCmd = &cobra.Command{
	Use:   "goprompt",
	Short: "GoPrompt is a customizable parallel prompt generator",
	Long:  `GoPrompt builds a prompt by running multiple segments in parallel, with a short timeout for each.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior if no subcommand is specified
		fmt.Println("[rootCmd] Default aggregator-based render. (TODO)")
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// GetRootCmd returns a reference to the root command (for testing).
func GetRootCmd() *cobra.Command {
	return rootCmd
}

func main() {
	Execute()
}
