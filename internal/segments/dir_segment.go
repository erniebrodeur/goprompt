package segments

import (
	"fmt"
	"os"
	"strings"
)

type DirSegment struct {
	ShowComponents int
}

func (d *DirSegment) Render(theme map[string]string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "[ERR]", err
	}
	comps := strings.Split(path, "/")
	if d.ShowComponents > 0 && len(comps) > d.ShowComponents {
		comps = comps[len(comps)-d.ShowComponents:]
	}
	dir := strings.Join(comps, "/")

	// Minimal color approach
	colorKey := "dir.normal"
	if theme != nil {
		hexColor, ok := theme[colorKey]
		if !ok {
			hexColor = "#FFFFFF"
		}
		ansi := fmt.Sprintf("\033[38;2;255;255;255m") // fallback
		// In reality, parse hexColor => r,g,b -> build ansi code
		dir = ansi + dir + "\033[0m"
	}

	return dir, nil
}

