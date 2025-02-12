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

// We'll define a subcommand that runs aggregator with a couple segments:
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Render the prompt with optional overrides",
	Long:  `Invokes the aggregator for a quick parallel prompt. If needed, you can pass --theme or --layout to override environment defaults.`,
	Run: func(cmd *cobra.Command, args []string) {
		agg := aggregator.New(100 * time.Millisecond)
		agg.Segments = []segments.Segment{
			&segments.DirSegment{ShowComponents: 1},
			&segments.GitSegment{},
		}

		out := agg.Collect(nil) // ignoring themeMap for now
		fmt.Println(out)
	},
}

func init() {
	renderCmd.Flags().StringVarP(&themeOverride, "theme", "", "", "Override theme for a single run")
	renderCmd.Flags().StringVarP(&layoutOverride, "layout", "", "", "Override layout for a single run")
}

// Getter for tests
func GetRootCmd() *cobra.Command {
	return rootCmd
}

---
