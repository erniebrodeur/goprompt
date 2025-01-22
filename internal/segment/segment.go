package segment

import (
	"fmt"
	"strings"

	"github.com/erniebrodeur/goprompt/internal/model"
)

type Segment interface {
    Name() string
    Enabled() bool
    Output() (model.SegmentOutput, error)
}

func (m *Manager) BuildPrompt(width int, displayWidth func(string) int) string {
    leftStr := m.buildSideLeft(m.LeftSegments)
    rightStr := m.buildSideRight(m.RightSegments)
    leftLen := displayWidth(leftStr)
    rightLen := displayWidth(rightStr)
    gap := width - (leftLen + rightLen)
    if gap < 0 {
        gap = 0
    }
    filler := strings.Repeat("─", gap)
    return leftStr + filler + rightStr
}

func (m *Manager) buildSideLeft(segs []Segment) string {
    out := ""
    i := 0
    for i < len(segs) {
        seg := segs[i]
        if !seg.Enabled() {
            i++
            continue
        }
        data, err := seg.Output()
        if err != nil || data.Text == "" {
            i++
            continue
        }
        if seg.Name() == "directory" && (i+1) < len(segs) {
            nextSeg := segs[i+1]
            if nextSeg.Enabled() && nextSeg.Name() == "git" {
                nextData, _ := nextSeg.Output()
                if nextData.Text != "" {
                    data.Text = data.Text + ":" + nextData.Text
                }
                i += 2
            } else {
                i++
            }
        } else {
            i++
        }
        colored := m.Theme.Colorize(data)
        if out == "" {
            out += fmt.Sprintf("─┤ %s ├──", colored)
        } else {
            out += fmt.Sprintf("┤ %s ├", colored)
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
