package aggregator

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/erniebrodeur/goprompt/internal/segments"
)

type mockSegment struct {
	result string
	delay  time.Duration
	err    error
}

func (m *mockSegment) Render(theme map[string]string) (string, error) {
	if m.delay > 0 {
		time.Sleep(m.delay)
	}
	if m.err != nil {
		return "[ERR]", m.err
	}
	return m.result, nil
}

func TestAggregator(t *testing.T) {
	tests := []struct {
		name     string
		segments []segments.Segment
		timeout  time.Duration
		want     string
	}{
		{
			name: "AllSuccess",
			segments: []segments.Segment{
				&mockSegment{result: "DIR"},
				&mockSegment{result: "GIT"},
			},
			timeout: 100 * time.Millisecond,
			want:    "DIR GIT",
		},
		{
			name: "SingleError",
			segments: []segments.Segment{
				&mockSegment{result: "OK"},
				&mockSegment{err: errors.New("some failure")},
			},
			timeout: 100 * time.Millisecond,
			want:    "OK [ERR]",
		},
		{
			name: "SlowSegmentTimesOut",
			segments: []segments.Segment{
				&mockSegment{result: "FAST", delay: 10 * time.Millisecond},
				&mockSegment{result: "SLOW", delay: 200 * time.Millisecond},
			},
			timeout: 50 * time.Millisecond,
			want:    "FAST [ERR]",
		},
		{
			name:     "NoSegments",
			segments: []segments.Segment{},
			timeout:  100 * time.Millisecond,
			want:     "[No Segments Configured]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agg := New(tt.timeout)
			agg.Segments = tt.segments

			got := agg.Collect(nil)
			if got != tt.want {
				t.Errorf("%s: Collect() = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

func TestAggregatorPartialWait(t *testing.T) {
	agg := New(80 * time.Millisecond)
	agg.Segments = []segments.Segment{
		&mockSegment{result: "A", delay: 10 * time.Millisecond},
		&mockSegment{result: "B", delay: 70 * time.Millisecond},
		&mockSegment{result: "C", delay: 200 * time.Millisecond},
	}
	got := agg.Collect(nil)
	if !strings.HasPrefix(got, "A B") {
		t.Errorf("Expected 'A B ...', got %q", got)
	}
	if !strings.Contains(got, "[ERR]") {
		t.Errorf("Expected third segment to be [ERR], got %q", got)
	}
}

func TestAggregatorContextCancel(t *testing.T) {
	agg := New(200 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	results := make([]string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	agg.Segments = []segments.Segment{
		&mockSegment{result: "X", delay: 50 * time.Millisecond},
		&mockSegment{result: "Y", delay: 10 * time.Millisecond},
	}

	go func() {
		defer wg.Done()
		got := agg.Collect(nil)
		results[0] = got
	}()

	select {
	case <-ctx.Done():
		// aggregator didn't finish in 20ms
		wg.Done()
		results[1] = "[CTX_CANCELED]"
	}

	wg.Wait()

	if results[1] == "[CTX_CANCELED]" && !strings.Contains(results[0], "[ERR]") {
		t.Errorf("Expected aggregator partial result with [ERR], got %q", results[0])
	}
}
