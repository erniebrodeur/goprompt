package aggregator

import (
	"context"
	"time"
)

// Segment is a simple interface for segments that produce a portion of the prompt.
type Segment interface {
	Output(ctx context.Context) (string, error)
}

// BuildPrompt is a minimal placeholder that compiles but doesn't yet implement the logic.
// You'll expand it to pass the BDD scenarios in aggregator_fallback.feature / aggregator_godog_test.go.
func BuildPrompt(ctx context.Context, segs []Segment, timeout time.Duration) (string, error) {
	return "", nil
}

