package segment

import (
	"fmt"
	"time"

	"github.com/erniebrodeur/goprompt/internal/colors"
)

type TimeSegment struct{}

func NewTimeSegment() *TimeSegment {
    return &TimeSegment{}
}

func (t *TimeSegment) Name() string { return "time" }

func (t *TimeSegment) Enabled() bool { return true }

func (t *TimeSegment) Render() (string, error) {
    now := time.Now().Format("15:04:05")
    return fmt.Sprintf("%s%s%s", colors.Green, now, colors.Reset), nil
}
