package goprompt

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goprompt",
	Short: "GoPrompt is a customizable parallel prompt generator",
	Long:  `GoPrompt builds a prompt by running multiple segments in parallel, with a short timeout for each.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is specified, default to "render" logic or similar
		fmt.Println("[rootCmd] Defaulting to aggregator-based render. (TODO)")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Typically add subcommands here, e.g.:
	// rootCmd.AddCommand(renderCmd)
	// rootCmd.AddCommand(themeCmd)
	// etc.
}
