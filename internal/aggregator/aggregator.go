package aggregator

import (
	"time"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

type Aggregator struct {
	Timeout  time.Duration
	Segments []segments.Segment
}

func New(t time.Duration) *Aggregator {
	return &Aggregator{
		Timeout: t,
	}
}

func (a *Aggregator) Collect(themeMap map[string]string) string {
	// Placeholder implementation
	return "NOT IMPLEMENTED"
}

