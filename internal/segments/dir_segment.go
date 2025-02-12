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
	if theme != nil {
		if hexColor, ok := theme["dir.normal"]; ok {
			// In reality, parse hexColor => r,g,b => build ANSI
			// We'll do a placeholder:
			ansi := "\033[38;2;255;255;255m"
			dir = ansi + dir + "\033[0m"
			_ = hexColor
		}
	}

	return dir, nil
}
