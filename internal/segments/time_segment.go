package segments

import "fmt"

type TimeSegment struct {
Format string
}

func (t *TimeSegment) Render(theme map[string]string) (string, error) {
// Minimal placeholder so it compiles
// For a real solution, parse t.Format or default to e.g. "15:04"
return fmt.Sprintf("TIME(%s)", t.Format), nil
}

