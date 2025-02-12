package segments

import (
	"fmt"
	"os"
	"strings"
)

type DirSegment struct {
	// Possibly store any config, e.g. how many path components to show
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

	// If theme is set, color it:
	colorKey := "dir.normal"
	hexColor, ok := theme[colorKey]
	if !ok {
		hexColor = "#FFFFFF" // fallback
	}
	colorPrefix := fmt.Sprintf("\033[38;2;%s;%s;%sm", "255", "255", "255") // TODO parse hexColor properly

	// Minimal version - parseHexColor is in theme package, we skip for brevity
	return colorPrefix + dir + "\033[0m", nil
}
