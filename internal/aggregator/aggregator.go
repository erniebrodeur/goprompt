package aggregator

import (
	"context"
	"time"
)

type Segment interface {
	Output(ctx context.Context) (string, error)
}

func BuildPrompt(ctx context.Context, segs []Segment, timeout time.Duration) (string, error) {
	return "", nil
}

