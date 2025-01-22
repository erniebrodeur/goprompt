package theme

import "github.com/erniebrodeur/goprompt/internal/segment"

// Theme applies color and formatting to each segment’s raw output.
type Theme interface {
    Colorize(segData segment.SegmentOutput) string
}
