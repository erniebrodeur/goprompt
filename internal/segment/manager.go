package segment

import (
	"fmt"
	"os"
	"os/user"
)

// Manager orchestrates segment rendering. It gathers shared data into a Context,
// then calls each segment’s Output(...) method with that context.
type Manager struct {
	LeftSegments  []Segment
	RightSegments []Segment
	Theme         Theme

	ctx *Context // Holds environment info to pass to segments
}

// BuildPrompt constructs the final prompt string using the specified width
// and a function to measure display width (e.g., removing ANSI codes).
func (m *Manager) BuildPrompt(width int, displayWidthFn func(string) int) string {
	m.initContext() // Gather environment data once

	leftStr := m.buildSegments(m.LeftSegments)
	rightStr := m.buildSegments(m.RightSegments)

	// Here, you can do padding/fillers between leftStr & rightStr if you want.
	// For simplicity, just return left + right.
	return fmt.Sprintf("%s%s", leftStr, rightStr)
}

// initContext populates the Manager’s Context with data like the current directory or root status.
func (m *Manager) initContext() {
	m.ctx = &Context{}

	// Get current working directory
	if wd, err := os.Getwd(); err == nil {
		m.ctx.Pwd = wd
	}

	// Check if the user is root
	if u, err := user.Current(); err == nil && u.Uid == "0" {
		m.ctx.IsRoot = true
	}

	// If you want to detect SSH, set m.ctx.IsSSH = true if certain environment vars are set, etc.
}

// buildSegments calls each segment’s Output method with the shared context.
func (m *Manager) buildSegments(segs []Segment) string {
	out := ""
	for _, s := range segs {
		if s.Enabled(m.ctx) {
			res, err := s.Output(m.ctx)
			if err != nil {
				// Optionally handle or log error
				continue
			}
			colored := m.Theme.Colorize(res)
			// If you want a space or separator between segments, add here
			out += colored + " "
		}
	}
	return out
}
