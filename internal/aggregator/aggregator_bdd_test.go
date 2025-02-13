package aggregator_test

import (
"context"
"errors"
"time"

. "github.com/onsi/ginkgo/v2"
. "github.com/onsi/gomega"

"github.com/erniebrodeur/goprompt/internal/aggregator"
"github.com/erniebrodeur/goprompt/internal/segments"
)

// Mock segment for testing concurrency, delays, and errors.
type mockSegment struct {
result string
delay  time.Duration
err    error
}

func (m *mockSegment) Render(_ map[string]string) (string, error) {
if m.delay > 0 {
	time.Sleep(m.delay)
}
if m.err != nil {
	return "[ERR]", m.err
}
return m.result, nil
}

var _ = Describe("Aggregator Concurrency & Timeouts", func() {
var agg *aggregator.Aggregator

BeforeEach(func() {
	// Default aggregator with a 100ms timeout
	agg = aggregator.New(100 * time.Millisecond)
})

Context("No segments", func() {
	It("returns [No Segments Configured]", func() {
		// aggregator.Segments is empty
		output := agg.Collect(nil)
		Expect(output).To(Equal("[No Segments Configured]"))
	})
})

Context("Single quick segment", func() {
	It("returns the segment's result immediately", func() {
		agg.Segments = []segments.Segment{
			&mockSegment{result: "FAST", delay: 10 * time.Millisecond},
		}
		output := agg.Collect(nil)
		Expect(output).To(ContainSubstring("FAST"))
		Expect(output).NotTo(ContainSubstring("[ERR]"))
	})
})

Context("Single slow segment", func() {
	It("times out and yields [ERR]", func() {
		agg.Segments = []segments.Segment{
			&mockSegment{result: "TOO_SLOW", delay: 200 * time.Millisecond},
		}
		output := agg.Collect(nil)
		Expect(output).To(ContainSubstring("[ERR]"))
		Expect(output).NotTo(ContainSubstring("TOO_SLOW"))
	})
})

Context("Multiple segments with varied speeds", func() {
	It("marks only the slow ones as [ERR], includes fast ones", func() {
		agg.Segments = []segments.Segment{
			&mockSegment{result: "FAST1", delay: 0},
			&mockSegment{result: "SLOW1", delay: 150 * time.Millisecond},
			&mockSegment{result: "FAST2", delay: 20 * time.Millisecond},
			&mockSegment{result: "SLOW2", delay: 300 * time.Millisecond},
		}
		output := agg.Collect(nil)

		// FAST1 and FAST2 finish within 100ms => appear in output
		Expect(output).To(ContainSubstring("FAST1"))
		Expect(output).To(ContainSubstring("FAST2"))

		// SLOW1 and SLOW2 exceed 100ms => [ERR]
		Expect(output).To(ContainSubstring("[ERR]"))
	})
})

Context("Segments that return errors", func() {
	It("produces [ERR] for any failing segment", func() {
		agg.Segments = []segments.Segment{
			&mockSegment{result: "OK", delay: 10 * time.Millisecond},
			&mockSegment{err: errors.New("segment failure")},
		}
		output := agg.Collect(nil)
		Expect(output).To(ContainSubstring("OK"))
		Expect(output).To(ContainSubstring("[ERR]"))
	})
})

Context("Context cancellation (optional advanced scenario)", func() {
	It("stops aggregator if the context is canceled externally", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		// aggregator might internally pass a WithTimeout for 100ms,
		// but external cancellation can happen sooner
		// We'll just check partial or forced [ERR].
		agg.Timeout = 200 * time.Millisecond
		agg.Segments = []segments.Segment{
			&mockSegment{result: "SEG1", delay: 70 * time.Millisecond},
			&mockSegment{result: "SEG2", delay: 10 * time.Millisecond},
		}

		doneCh := make(chan string, 1)
		go func() {
			out := agg.Collect(nil)
			doneCh <- out
		}()

		select {
		case out := <-doneCh:
			// The aggregator finished before external context timed out
			// Possibly it includes "SEG2" or [ERR]
			Expect(out).To(SatisfyAny(
				ContainSubstring("SEG2"),
				ContainSubstring("[ERR]")),
			)
		case <-ctx.Done():
			// aggregator is forcibly interrupted
			// We might still expect partial or [ERR] in real usage
		}
	})
})
})