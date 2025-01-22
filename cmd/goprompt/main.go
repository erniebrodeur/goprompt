package main

import (
	"fmt"
	"os"

	"github.com/erniebrodeur/goprompt/internal/segment"
	"github.com/erniebrodeur/goprompt/internal/term"
	"github.com/erniebrodeur/goprompt/internal/theme"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	width, err := term.GetWidth()
	if err != nil {
		return err
	}

	// Create a Manager that has some segments on the left and right
	mgr := &segment.Manager{
		LeftSegments: []segment.Segment{
			segment.NewDirSegment(),
			segment.NewUserSegment(),
			// segment.NewGitSegment(), etc.
		},
		RightSegments: []segment.Segment{
			// segment.NewTimeSegment(),
		},
		Theme: theme.NewDefaultTheme(),
	}

	prompt := mgr.BuildPrompt(width, term.DisplayWidth)
	fmt.Print(prompt)
	return nil
}
