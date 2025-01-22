package main

import (
	"fmt"
	"os"

	"github.com/erniebrodeur/goprompt/internal/segment"
	"github.com/erniebrodeur/goprompt/internal/term"
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
        return fmt.Errorf("failed to get terminal width: %w", err)
    }

    mgr := &segment.Manager{
        LeftSegments: []segment.Segment{
            // Add as many left segments as you want:
            segment.NewDirSegment(),
            segment.NewUserSegment(),
            // e.g. segment.NewGitSegment(),
        },
        RightSegments: []segment.Segment{
            // Put your right-side segments here:
            segment.NewTimeSegment(),
        },
    }

    prompt := mgr.BuildPrompt(width, term.DisplayWidth)
    fmt.Print(prompt)
    return nil
}
