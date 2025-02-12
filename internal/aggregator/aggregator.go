package aggregator

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

type Aggregator struct {
	Segments []segments.Segment
	Timeout  time.Duration
}

func New(timeout time.Duration) *Aggregator {
	return &Aggregator{
		Timeout: timeout,
	}
}

// Collect runs each Segment in parallel under a short timeout,
// returning a single prompt string.
func (a *Aggregator) Collect(themeMap map[string]string) string {
	if len(a.Segments) == 0 {
		return "[No Segments Configured]"
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.Timeout)
	defer cancel()

	results := make([]string, len(a.Segments))

	var wg sync.WaitGroup
	wg.Add(len(a.Segments))

	for i, seg := range a.Segments {
		go func(idx int, s segments.Segment) {
			defer wg.Done()
			out, err := s.Render(themeMap)
			if err != nil {
				results[idx] = "[ERR]"
				return
			}
			results[idx] = out
		}(i, seg)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// All segments finished in time
	case <-ctx.Done():
		// Timeout, mark unfinished segments as [ERR]
		for i, v := range results {
			if v == "" {
				results[i] = "[ERR]"
			}
		}
	}

	return strings.Join(results, " ")
}
