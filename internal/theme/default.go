package theme

import (
	"github.com/erniebrodeur/goprompt/internal/model"
	"github.com/muesli/termenv"
)

// DefaultTheme uses termenv for 24-bit (or best available) color.
type DefaultTheme struct {
	profile termenv.Profile
}

// NewDefaultTheme creates a DefaultTheme, storing a single color profile.
func NewDefaultTheme() *DefaultTheme {
	return &DefaultTheme{
		profile: termenv.EnvColorProfile(), // Detect once
	}
}

// Colorize decides which color to apply based on segment name and flags.
func (d *DefaultTheme) Colorize(segData model.SegmentOutput) string {
	switch segData.Name {
	case "directory":
		// Example: bright blue
		style := termenv.Style{}.Foreground(d.profile.Color("#005fff"))
		return style.Styled(segData.Text)

	case "user":
		// Red for root, green for normal user
		if segData.IsRoot {
			style := termenv.Style{}.Foreground(d.profile.Color("#ff0000"))
			return style.Styled(segData.Text)
		}
		style := termenv.Style{}.Foreground(d.profile.Color("#00ff00"))
		return style.Styled(segData.Text)

	case "time":
		// Cyan
		style := termenv.Style{}.Foreground(d.profile.Color("#00ffff"))
		return style.Styled(segData.Text)

	case "git":
		// Yellow
		style := termenv.Style{}.Foreground(d.profile.Color("#ffff00"))
		return style.Styled(segData.Text)

	default:
		// No color for unrecognized segment
		return segData.Text
	}
}
