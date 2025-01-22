package theme

import "github.com/erniebrodeur/goprompt/internal/model"

// Theme applies color and formatting to each segment’s raw output.
type Theme interface {
    Colorize(segData model.SegmentOutput) string
}
