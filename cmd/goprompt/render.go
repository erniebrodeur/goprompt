package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/erniebrodeur/goprompt/internal/aggregator"
	"github.com/erniebrodeur/goprompt/internal/segments"
)

var (
	themeOverride  string
	layoutOverride string
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render the prompt with optional overrides",
	Long: `render subcommand triggers the aggregator with minimal defaults.
You could parse --theme or --layout flags to override environment-based defaults.`,
	Run: func(cmd *cobra.Command, args []string) {
		// You might apply themeOverride or layoutOverride, but for now we’ll ignore them.

		// Create aggregator with a 100ms timeout
		agg := aggregator.New(100 * time.Millisecond)

		// Add some example segments
		agg.Segments = []segments.Segment{
			&segments.DirSegment{ShowComponents: 1},
			&segments.GitSegment{},
			// Additional segments can be added here
		}

		// For theming, we might pass a real map from environment or override flags
		finalOutput := agg.Collect(nil)
		fmt.Println(finalOutput)
	},
}

func init() {
	renderCmd.Flags().StringVarP(&themeOverride, "theme", "", "", "Override theme for a single run")
	renderCmd.Flags().StringVarP(&layoutOverride, "layout", "", "", "Override layout for a single run")
}

