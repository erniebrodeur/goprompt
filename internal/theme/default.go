package theme

import (
	"github.com/erniebrodeur/goprompt/internal/model"
	"github.com/muesli/termenv"
)

// DefaultTheme uses termenv for 24-bit color
type DefaultTheme struct{}

// Colorize applies different colors based on segment name or flags (e.g., IsRoot).
func (d DefaultTheme) Colorize(segData model.SegmentOutput) string {
    // Determine terminal color profile once
    p := termenv.ColorProfile()

    // Base text
    styled := termenv.String(segData.Text)

    switch segData.Name {
    case "directory":
        // Example: bright blue
        styled = styled.Foreground(p.Color("#005fff"))

    case "user":
        if segData.IsRoot {
            styled = styled.Foreground(p.Color("#ff0000")) // bright red for root
        } else {
            styled = styled.Foreground(p.Color("#00ff00")) // green for normal user
        }

    case "time":
        styled = styled.Foreground(p.Color("#00ffff")) // cyan

    case "git":
        styled = styled.Foreground(p.Color("#ffff00")) // yellow

    // Add other segments or defaults as needed
    default:
        // No special color
    }

    return styled.String()
}
