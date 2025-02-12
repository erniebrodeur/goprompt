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
	// If you want real theming, parse the hex color, build an ANSI code, etc.
	if theme != nil {
		if hexColor, ok := theme["dir.normal"]; ok {
			ansi := "\033[38;2;255;255;255m" // Pretend parse
			_ = hexColor
			dir = ansi + dir + "\033[0m"
		}
	}

	return dir, nil
}

