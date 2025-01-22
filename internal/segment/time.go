package segment

import "time"

type TimeSegment struct{}

func NewTimeSegment() *TimeSegment {
    return &TimeSegment{}
}

func (t *TimeSegment) Name() string   { return "time" }
func (t *TimeSegment) Enabled() bool  { return true }

func (t *TimeSegment) Output() (SegmentOutput, error) {
    now := time.Now().Format("03:04pm ─ 1/2")
    return SegmentOutput{
        Name: "time",
        Text: now,
    }, nil
}
