package aggregator

import (
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

func TestAggregatorCollect(t *testing.T) {
	agg := New(50 * time.Millisecond)
	agg.Segments = []segments.Segment{
		&mockSegment{result: "DIR", delay: 10 * time.Millisecond},
		&mockSegment{result: "GIT", delay: 10 * time.Millisecond},
	}

	out := agg.Collect(nil)
	if out != "DIR GIT" {
		t.Errorf("expected 'DIR GIT', got '%s'", out)
	}
}

func TestAggregatorTimeout(t *testing.T) {
	agg := New(10 * time.Millisecond)
	agg.Segments = []segments.Segment{
		&mockSegment{result: "SLOW", delay: 50 * time.Millisecond},
	}

	out := agg.Collect(nil)
	if out != "[ERR]" {
		t.Errorf("expected '[ERR]', got '%s'", out)
	}
}
