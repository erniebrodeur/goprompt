package segment

import (
	"fmt"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/theme"
)

// Segment defines the minimal interface for a prompt segment.
// No color or glyphs should be returned—just raw data/flags.
type Segment interface {
    Name() string
    Enabled() bool
    Output() (SegmentOutput, error)
}

type Manager struct {
    LeftSegments  []Segment
    RightSegments []Segment
    // Theme to apply color/wrapping
    Theme theme.Theme
}

// BuildPrompt assembles the full line, including filler "─" between left & right.
func (m *Manager) BuildPrompt(width int, displayWidth func(string) int) string {
    leftStr := m.buildSideLeft(m.LeftSegments)
    rightStr := m.buildSideRight(m.RightSegments)

    leftLen := displayWidth(leftStr)
    rightLen := displayWidth(rightStr)
    gap := width - (leftLen + rightLen)
    if gap < 0 {
        gap = 0
    }

    // Insert "┤" plus a run of "─" to fill
    filler := "┤" + strings.Repeat("─", gap)
    return leftStr + filler + rightStr
}

func (m *Manager) buildSideLeft(segs []Segment) string {
    out := ""
    for i, seg := range segs {
        if seg.Enabled() {
            data, err := seg.Output()
            if err == nil && data.Text != "" {
                // 1) Theme color
                colored := m.Theme.Colorize(data)
                // 2) Wrap with bookshelf glyphs
                // first left segment => "─┤ <colored> ├──"
                // subsequent => "┤ <colored> ├"
                if i == 0 {
                    out += fmt.Sprintf("─┤ %s ├──", colored)
                } else {
                    out += fmt.Sprintf("┤ %s ├", colored)
                }
            }
        }
    }
    return out
}

func (m *Manager) buildSideRight(segs []Segment) string {
    out := ""
    for i, seg := range segs {
        if seg.Enabled() {
            data, err := seg.Output()
            if err == nil && data.Text != "" {
                colored := m.Theme.Colorize(data)
                // last right segment => "┤ <colored> ├─"
                // others => "┤ <colored> ├"
                if i == len(segs)-1 {
                    out += fmt.Sprintf("┤ %s ├─", colored)
                } else {
                    out += fmt.Sprintf("┤ %s ├", colored)
                }
            }
        }
    }
    return out
}
