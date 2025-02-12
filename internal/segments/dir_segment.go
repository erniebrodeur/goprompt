package segments

import (
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

	// If theme is provided, we could parse "dir.normal" and build an ANSI code, skipping for brevity
	return dir, nil
}

---
