package segment

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
    leftStr := m.buildSide(m.LeftSegments)
    rightStr := m.buildSide(m.RightSegments)

    leftLen := displayWidth(leftStr)
    rightLen := displayWidth(rightStr)

    gap := width - (leftLen + rightLen)
    if gap < 1 {
        gap = 1
    }

    return leftStr + makeSpaces(gap) + rightStr
}

func (m *Manager) buildSide(segs []Segment) string {
    out := ""
    for _, seg := range segs {
        if seg.Enabled() {
            txt, err := seg.Render()
            if err != nil {
                // Skip if there's an error in the segment
                continue
            }
            out += txt
        }
    }
    return out
}

func makeSpaces(n int) string {
    // Simple approach: strings.Repeat
    // But here is a direct snippet:
    return string(make([]byte, n))
}
