package segment

import (
	"fmt"
	"strings"
)

type Segment interface {
    Name() string
    Enabled() bool
    Render() (string, error)
}

type Manager struct {
    LeftSegments  []Segment
    RightSegments []Segment
}

func (m *Manager) BuildPrompt(width int, displayWidth func(string) int) string {
    leftStr  := m.buildSideLeft(m.LeftSegments)
    rightStr := m.buildSideRight(m.RightSegments)

    leftLen  := displayWidth(leftStr)
    rightLen := displayWidth(rightStr)
    gap := width - (leftLen + rightLen)
    if gap < 0 {
        gap = 0
    }

    // The filler: a "┤" plus a run of "─" dashes
    filler := "┤" + strings.Repeat("─", gap)

    // Combine them into the final single-line prompt
    return leftStr + filler + rightStr
}

// buildSideLeft: handles your left segments
func (m *Manager) buildSideLeft(segs []Segment) string {
    out := ""
    for i, seg := range segs {
        if seg.Enabled() {
            txt, err := seg.Render()
            if err == nil && txt != "" {
                switch i {
                case 0:
                    // The *first* left segment: "─┤ <txt> ├──"
                    out += wrapSegmentLeftFirst(txt)
                default:
                    // Subsequent left segments: "┤ <txt> ├"
                    out += wrapSegmentLeftNext(txt)
                }
            }
        }
    }
    return out
}

// buildSideRight: handles your right segments
func (m *Manager) buildSideRight(segs []Segment) string {
    out := ""
    for i, seg := range segs {
        if seg.Enabled() {
            txt, err := seg.Render()
            if err == nil && txt != "" {
                // If it's NOT the last one, or if you have multiple right segments:
                // "┤ <txt> ├"
                // If it's the final right segment: "┤ <txt> ├─"
                if i == len(segs)-1 {
                    out += wrapSegmentRightLast(txt)
                } else {
                    out += wrapSegmentRightMid(txt)
                }
            }
        }
    }
    return out
}

func wrapSegmentLeftFirst(s string) string {
    return fmt.Sprintf("─┤ %s ├──", s)
}

func wrapSegmentLeftNext(s string) string {
    return fmt.Sprintf("┤ %s ├", s)
}

func wrapSegmentRightMid(s string) string {
    return fmt.Sprintf("┤ %s ├", s)
}

func wrapSegmentRightLast(s string) string {
    return fmt.Sprintf("┤ %s ├─", s)
}
