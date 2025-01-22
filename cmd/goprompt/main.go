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

    // Ensure we assign a real Theme so it's never nil.
    mgr := &segment.Manager{
        LeftSegments: []segment.Segment{
            segment.NewDirSegment(),
            segment.NewUserSegment(),
            // segment.NewGitSegment(), etc. if needed
        },
        RightSegments: []segment.Segment{
            segment.NewTimeSegment(),
        },
        Theme: theme.DefaultTheme{}, // ← The crucial fix!
    }

    prompt := mgr.BuildPrompt(width, term.DisplayWidth)
    fmt.Print(prompt)
    return nil
}
