package theme

import (
	"fmt"

	"github.com/erniebrodeur/goprompt/internal/colors"
	"github.com/erniebrodeur/goprompt/internal/segment"
)

// DefaultTheme applies colors based on segment name or flags like IsRoot.
type DefaultTheme struct{}

func (d DefaultTheme) Colorize(segData segment.SegmentOutput) string {
    switch segData.Name {
    case "directory":
        return fmt.Sprintf("%s%s%s", colors.BrightBlue, segData.Text, colors.Reset)
    case "user":
        if segData.IsRoot {
            return fmt.Sprintf("%s%s%s", colors.BrightRed, segData.Text, colors.Reset)
        }
        return fmt.Sprintf("%s%s%s", colors.BrightGreen, segData.Text, colors.Reset)
    case "time":
        return fmt.Sprintf("%s%s%s", colors.BrightCyan, segData.Text, colors.Reset)
    case "git":
        return fmt.Sprintf("%s%s%s", colors.BrightYellow, segData.Text, colors.Reset)
    default:
        // If no match, just return the text uncolored
        return segData.Text
    }
}
