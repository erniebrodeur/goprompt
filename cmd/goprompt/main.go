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
	Long:  `GoPrompt builds a prompt by running multiple segments in parallel, each with a short timeout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is specified, we do a placeholder or you can call aggregator directly.
		fmt.Println("[rootCmd] Default aggregator-based render. (Try 'goprompt render' or other subcommands)")
	},
}

// Execute is the entry point from main().
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	// Add subcommands (e.g., render) in init()
	Execute()
}

func init() {
	// We'll add the 'renderCmd' from render.go here.
	rootCmd.AddCommand(renderCmd)
}

