package goprompt

import "github.com/spf13/cobra"

// If you want to expose the root command to tests, here's a simple getter:
func GetRootCmd() *cobra.Command {
	return rootCmd
}