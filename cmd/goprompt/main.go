package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goprompt",
	Short: "GoPrompt is a customizable parallel prompt generator",
	Long:  `GoPrompt builds a prompt by running multiple segments in parallel, each with a short timeout.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, we do a placeholder or aggregator logic
		fmt.Println("[rootCmd] Default aggregator-based render. (Try 'goprompt render')")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}

func init() {
	rootCmd.AddCommand(renderCmd)
}

---
