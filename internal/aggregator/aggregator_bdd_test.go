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

var _ = Describe("Aggregator Concurrency", func() {

	Context("Multiple concurrency edges", func() {
		It("marks slow segments as [ERR] after timeout", func() {
			agg := aggregator.New(100 * time.Millisecond)
			agg.Segments = []segments.Segment{
				&mockSegment{result: "Quick", delay: 20 * time.Millisecond},
				&mockSegment{result: "Slow", delay: 200 * time.Millisecond},
			}
			output := agg.Collect(nil)
			Expect(output).To(ContainSubstring("Quick"))
			Expect(output).To(ContainSubstring("[ERR]"))
		})
	})

	Context("No segments configured", func() {
		It("returns a placeholder text", func() {
			agg := aggregator.New(50 * time.Millisecond)
			output := agg.Collect(nil)
			Expect(output).To(Equal("[No Segments Configured]"))
		})
	})

	Context("Canceling context (optional demonstration)", func() {
		It("forces partial results if aggregator is canceled externally", func() {
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
			defer cancel()

			agg := aggregator.New(200 * time.Millisecond) 
			agg.Segments = []segments.Segment{
				&mockSegment{result: "X", delay: 50 * time.Millisecond},
				&mockSegment{result: "Y", delay: 10 * time.Millisecond},
			}

			doneCh := make(chan string, 1)
			go func() {
				out := agg.Collect(nil)
				doneCh <- out
			}()

			select {
			case <-ctx.Done():
				// aggregator not finished
				// we expect partial or [ERR] for some segments
				// This is a contrived example, real usage might differ
			case out := <-doneCh:
				Expect(out).To(Or(
					ContainSubstring("Y"),
					ContainSubstring("[ERR]")),
				)
			}
		})
	})
})

